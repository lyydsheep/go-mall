package router

// 路由注册相关的, 都放在router 包下

import (
	"github.com/faiz/go-mall/common/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(s *gin.Engine) {
	g := s.Group("/")
	g.Use(middleware.StartTrace(), middleware.LogAccess(), middleware.PanicRecorder())
	registerBuild(g)
}
