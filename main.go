package main

import (
	"github.com/faiz/go-mall/common/logger"
	_ "github.com/faiz/go-mall/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/readConfig", func(c *gin.Context) {
		logger.ZapLoggerTest()
	})
	r.Run()
}
