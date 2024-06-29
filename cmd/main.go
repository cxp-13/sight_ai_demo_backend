package main

import (
	"gateway_service/config"
	"gateway_service/internal/eth"
	"gateway_service/internal/handlers"
	"gateway_service/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.LoadConfig()
	eth.Init()
	go services.ListenOnChain()
	r.POST("/compute", handlers.Compute)
	r.Run(":8080")
}
