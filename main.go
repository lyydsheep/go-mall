package main

import (
	"github.com/faiz/go-mall/api/router"
	"github.com/faiz/go-mall/common/enum"
	"github.com/faiz/go-mall/config"
	_ "github.com/faiz/go-mall/config"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	router.RegisterRouters(g)
	if err := g.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
