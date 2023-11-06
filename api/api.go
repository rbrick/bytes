package api

import "github.com/gin-gonic/gin"

func NewServer() *gin.Engine {
	engine := gin.Default()

	v1Group := engine.Group("/api/v1")
	{
		v1Group.GET("/price", Price)
		v1Group.GET("/wallet/:address", Wallet)
	}
	return engine
}
