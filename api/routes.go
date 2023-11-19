package api

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rbrick/bytes/cache"
	"github.com/rbrick/bytes/contracts"
	"github.com/rbrick/bytes/models"
	"github.com/rbrick/bytes/price"
	"github.com/rbrick/bytes/util"
	"github.com/rbrick/bytes/wallet"
)

var (
	ADDRESS_PATTERN = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
)

type Routes interface {
	MarketStats(*gin.Context)
	Wallet(*gin.Context)
}

type BytesPrices struct {
	EtherValue *big.Float
	USDValue   *big.Float
}

type RoutesHandler struct {
	rpc            *ethclient.Client
	bytesETHPricer price.Pricer
	bytesUSDPricer price.Pricer

	bytesPrice *cache.LazyCachedValue[*BytesPrices]

	totalSupply *cache.LazyCachedValue[*big.Int]
}

// Price fetches and returns the price of BYTES in both Ether and USDC
func (rh *RoutesHandler) MarketStats(ctx *gin.Context) {

	bytesPrice, err := rh.bytesPrice.Value()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	supply, err := rh.totalSupply.Value()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	totalSupplyBytes := price.EtherConverter(supply)

	marketCap := new(big.Float).Mul(totalSupplyBytes, bytesPrice.USDValue)

	ctx.JSON(http.StatusOK, models.MarketStats{
		Price: models.Prices{
			USDPrice:   bytesPrice.USDValue.Text('g', 16),
			EtherPrice: bytesPrice.EtherValue.Text('g', 16),
		},
		TotalSupply: totalSupplyBytes.Text('g', 16),
		MarketCap:   marketCap.Text('g', 16),
	})

}

// Wallet fetches information about a passed in wallet such as citizens staked, BYTES present in the wallet, etc.
func (rh *RoutesHandler) Wallet(ctx *gin.Context) {
	addr := ctx.Param("address")

	if !ADDRESS_PATTERN.MatchString(addr) {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid address"))
		return
	}

	checksumAddr := common.HexToAddress(addr)

	stakerContract, err := contracts.NewNeoTokyoStaker(wallet.NeoTokyoStakerAddress, rh.rpc)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	erc20Contract, err := contracts.NewERC20(Bytes2Address, rh.rpc)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	positions, err := wallet.GetPositions(stakerContract, checksumAddr)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	pendingRewards, err := wallet.GetPendingRewards(stakerContract, checksumAddr)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	balance, err := erc20Contract.BalanceOf(nil, checksumAddr)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	totalStaked := big.NewFloat(0.0)
	var stakedCitizens []models.Citizen

	for _, position := range positions {
		totalStaked = new(big.Float).Add(totalStaked, position.StakedBytes)

		stakedCitizens = append(stakedCitizens, models.Citizen{
			ID:     position.Id,
			Season: position.Season,
			Image:  fmt.Sprintf("https://citgen.rcw.io/s%d/600x600/%d", position.Season, position.Id),
		})
	}

	walletResponse := &models.Wallet{
		TokenBalance:   util.ConvertHuman(balance, 18).Text('f', 16),
		TotalStake:     totalStaked.Text('f', 16),
		PendingRewards: util.ConvertHuman(pendingRewards, 18).Text('f', 16),
		StakedCitizens: stakedCitizens,
	}

	ctx.JSON(http.StatusOK, walletResponse)
}

// bytesComputer is a lazy computer for the price of 'BYTES'. The prices are cached in-memory
// for 5 minutes, and when it is expired, this function will be called to recompute the value.
// It isn't actively re-calculated i.e. it's calculated upon access.
func (rh *RoutesHandler) bytesComputer() (*BytesPrices, error) {
	priceInEther, err := rh.bytesETHPricer.Price()

	if err != nil {
		return nil, err
	}

	priceInUsd, err := rh.bytesUSDPricer.Price()

	if err != nil {
		return nil, err
	}

	convertedUSD := util.ConvertHuman(priceInUsd, 6)
	convertedEther := util.ConvertHuman(priceInEther, 18)

	return &BytesPrices{
		USDValue:   convertedUSD,
		EtherValue: convertedEther,
	}, nil
}

func NewRoutes(rpcClient *ethclient.Client) Routes {
	routesHandler := &RoutesHandler{
		rpc:            rpcClient,
		bytesETHPricer: price.NewTokenPricer(rpcClient, Bytes2Address, common.HexToAddress("0x0")),
		bytesUSDPricer: price.NewTokenPricer(rpcClient, Bytes2Address, USDCAddress),
	}
	routesHandler.bytesPrice = cache.NewLazilyCachedValue[*BytesPrices](5*time.Minute, routesHandler.bytesComputer)
	routesHandler.totalSupply = cache.NewLazilyCachedValue[*big.Int](5*time.Minute, func() (*big.Int, error) {
		erc20, err := contracts.NewERC20(Bytes2Address, rpcClient)

		if err != nil {
			return nil, err
		}

		return erc20.TotalSupply(nil)
	})
	return routesHandler
}
