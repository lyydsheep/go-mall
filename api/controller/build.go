package controller

// controller 包中文件放置handler 方法

import (
	"github.com/faiz/go-mall/common/app"
	"github.com/faiz/go-mall/common/errcode"
	"github.com/faiz/go-mall/logic/appService"
	"github.com/gin-gonic/gin"
)

type BuildController struct {
	appDemoService appService.DemoAppService
}

func NewBuildController(app appService.DemoAppService) *BuildController {
	return &BuildController{
		appDemoService: app,
	}
}

func (build *BuildController) TestPagination(c *gin.Context) {
	data := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{
			Name: "faiz",
			Age:  18,
		},
		{
			Name: "lyy",
			Age:  10,
		},
	}
	p := app.NewPagination(c)
	p.SetTotalRows(len(data))
	app.NewResponse(c).SetPagination(p).Success(data)
}

func (build *BuildController) TestGormLog(c *gin.Context) {
	ids, err := build.appDemoService.GetAllIdentities(c)
	if err != nil {
		app.NewResponse(c).Error(errcode.ErrServer.WithCause(err))
	}
	app.NewResponse(c).Success(ids)
}
