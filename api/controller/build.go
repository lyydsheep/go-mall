package controller

// controller 包中文件放置handler 方法

import (
	"github.com/faiz/go-mall/common/app"
	"github.com/gin-gonic/gin"
)

func TestPagination(c *gin.Context) {
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
