package router

// 路由注册相关的, 都放在router 包下

import (
	"github.com/faiz/go-mall/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutersAndMiddleware(build *controller.BuildController, fs ...gin.HandlerFunc) *gin.Engine {
	s := gin.Default()
	RegisterMiddleware(s, fs...)

	g := s.Group("/")
	registerBuild(g, build)
	return s
}

func RegisterMiddleware(g *gin.Engine, fs ...gin.HandlerFunc) *gin.Engine {
	g.Use(fs...)
	return g
}
