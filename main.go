package main

import (
	log "github.com/faiz/go-mall/common/logger"
	"github.com/faiz/go-mall/common/middleware"
	_ "github.com/faiz/go-mall/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger(), middleware.StartTrace())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/logTest", func(c *gin.Context) {
		log.New(c).Info("this is a test", "test", "begin")
		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})
	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
