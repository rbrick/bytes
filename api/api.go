package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	Bytes2Address = common.HexToAddress("0xa19f5264F7D7Be11c451C093D8f92592820Bea86")
	USDCAddress   = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}

func NewServer(rpcClient *ethclient.Client) *Server {
	engine := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))
	routesHandler := NewRoutes(rpcClient)

	v1Group := engine.Group("/v1/api")
	{
		v1Group.GET("/price", routesHandler.Price)
		v1Group.GET("/wallet/:address", routesHandler.Wallet)
	}

	return &Server{
		engine: engine,
	}
}
