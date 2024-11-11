package main

import (
	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusOK, gin.H{
			"type": config.DataBase.Type,
			"dsn":  config.DataBase.Dsn,
		})
	})
	server.Run(":8080")
}
