package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyydsheep/go-mall/common/logger"
	"github.com/lyydsheep/go-mall/config"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.GET("/config-read", func(ctx *gin.Context) {
		logger.ZapLoggerTest()
		ctx.JSON(http.StatusOK, gin.H{
			"type":     config.DataBase.Type,
			"max_life": config.DataBase.MaxFileTime,
		})
	})
	server.Run(":8080")
}
