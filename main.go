package main

import (
	"errors"
	"github.com/faiz/go-mall/common/errcode"
	log "github.com/faiz/go-mall/common/logger"
	"github.com/faiz/go-mall/common/middleware"
	_ "github.com/faiz/go-mall/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.StartTrace(), middleware.LogAccess(), middleware.PanicRecorder())
	r.GET("/testError", func(c *gin.Context) {
		err := errors.New("gorm error")
		aErr := errcode.Wrap("包装", err)
		bErr := errcode.Wrap("再包装", aErr)
		log.New(c).Error("记录错误", "err", bErr)
		domainErr := errors.New("domain Error")
		apiErr := errcode.ErrServer.WithCause(domainErr)
		log.New(c).Error("记录错误", "err", apiErr)
		c.JSON(apiErr.HttpStatusCode(), gin.H{
			"code":    apiErr.GetCode(),
			"message": apiErr.GetMsg(),
		})
	})
	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
