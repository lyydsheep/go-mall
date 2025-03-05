package main

import (
	"github.com/faiz/go-mall/common/app"
	"github.com/faiz/go-mall/common/middleware"
	_ "github.com/faiz/go-mall/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.StartTrace(), middleware.LogAccess(), middleware.PanicRecorder())
	r.GET("/pagination", func(c *gin.Context) {
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
	})
	err := r.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
