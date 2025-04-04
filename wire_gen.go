// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/faiz/go-mall/api/controller"
	"github.com/faiz/go-mall/api/router"
	"github.com/faiz/go-mall/common/middleware"
	"github.com/faiz/go-mall/dal/cache"
	"github.com/faiz/go-mall/dal/dao"
	"github.com/faiz/go-mall/logic/appService"
	"github.com/faiz/go-mall/logic/domainService"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitializeApp() *gin.Engine {
	demoDAOV1 := dao.NewDemoDAO()
	demoCacheV1 := cache.NewCacheV1()
	demoDomainServiceV1 := domainService.NewDemoDomainServiceV1(demoDAOV1, demoCacheV1)
	demoAppServiceV1 := appService.NewDemoAppServiceV1(demoDomainServiceV1)
	buildController := controller.NewBuildController(demoAppServiceV1)
	v := middleware.GetHandlerFunc()
	engine := router.RegisterRoutersAndMiddleware(buildController, v...)
	return engine
}
