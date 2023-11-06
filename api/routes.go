package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
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
}

// Price fetches and returns the price of BYTES in both Ether and USDC
func (rh *RoutesHandler) Price(ctx *gin.Context) {
	priceInEther, err := rh.bytesETHPricer.Price()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	priceInUsd, err := rh.bytesUSDPricer.Price()

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	convertedUSD := util.ConvertHuman(priceInUsd, 6)
	convertedEther := util.ConvertHuman(priceInEther, 18)

	ctx.JSON(http.StatusOK, models.Prices{
		USDPrice:   convertedUSD.Text('f', 16),
		EtherPrice: convertedEther.Text('f', 16),
	})

}

// Wallet fetches information about a passed in wallet such as citizens staked, BYTES present in the wallet, etc.
func (*RoutesHandler) Wallet(ctx *gin.Context) {}

func NewRoutes(rpcClient *ethclient.Client) Routes {

	return &RoutesHandler{
		bytesETHPricer: price.NewTokenPricer(rpcClient, Bytes2Address, common.HexToAddress("0x0")),
		bytesUSDPricer: price.NewTokenPricer(rpcClient, Bytes2Address, USDCAddress),
	}
}
