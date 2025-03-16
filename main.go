package main

import (
	"github.com/faiz/go-mall/common/enum"
	log "github.com/faiz/go-mall/common/logger"
	"github.com/faiz/go-mall/config"
	"github.com/faiz/go-mall/dal/cache"
	"github.com/faiz/go-mall/dal/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitConfig()
	log.InitLogger()
	cache.RedisInit()
	dao.InitGormLogger()
	dao.InitDB()
}

func main() {
	if config.App.Env == enum.ModePROD {
		gin.SetMode(gin.ReleaseMode)
	}
	g := InitializeApp()
	if err := g.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
