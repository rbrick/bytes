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
	Price(*gin.Context)
	Wallet(*gin.Context)
}

type RoutesHandler struct {
	rpc            *ethclient.Client
	bytesETHPricer price.Pricer
	bytesUSDPricer price.Pricer

	bytesPrice *cache.LazyCachedValue[models.Prices]
}

// Price fetches and returns the price of BYTES in both Ether and USDC
func (rh *RoutesHandler) Price(ctx *gin.Context) {
	price, err := rh.bytesPrice.Value()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, price)
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
func (rh *RoutesHandler) bytesComputer() (models.Prices, error) {
	priceInEther, err := rh.bytesETHPricer.Price()

	if err != nil {
		return models.Prices{}, err
	}

	priceInUsd, err := rh.bytesUSDPricer.Price()

	if err != nil {
		return models.Prices{}, err
	}

	convertedUSD := util.ConvertHuman(priceInUsd, 6)
	convertedEther := util.ConvertHuman(priceInEther, 18)

	return models.Prices{
		USDPrice:   convertedUSD.Text('f', 16),
		EtherPrice: convertedEther.Text('f', 16),
	}, nil
}

func NewRoutes(rpcClient *ethclient.Client) Routes {
	routesHandler := &RoutesHandler{
		rpc:            rpcClient,
		bytesETHPricer: price.NewTokenPricer(rpcClient, Bytes2Address, common.HexToAddress("0x0")),
		bytesUSDPricer: price.NewTokenPricer(rpcClient, Bytes2Address, USDCAddress),
	}
	routesHandler.bytesPrice = cache.NewLazilyCachedValue[models.Prices](5*time.Minute, routesHandler.bytesComputer)
	return routesHandler
}
