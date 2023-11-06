package api

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rbrick/bytes/cache"
	"github.com/rbrick/bytes/models"
	"github.com/rbrick/bytes/price"
	"github.com/rbrick/bytes/util"
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
func (*RoutesHandler) Wallet(ctx *gin.Context) {

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
		bytesETHPricer: price.NewTokenPricer(rpcClient, Bytes2Address, common.HexToAddress("0x0")),
		bytesUSDPricer: price.NewTokenPricer(rpcClient, Bytes2Address, USDCAddress),
	}
	routesHandler.bytesPrice = cache.NewLazilyCachedValue[models.Prices](5*time.Minute, routesHandler.bytesComputer)
	return routesHandler
}
