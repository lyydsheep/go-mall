package router

import (
	"github.com/faiz/go-mall/api/controller"
	"github.com/gin-gonic/gin"
)

func registerBuild(s *gin.RouterGroup) {
	g := s.Group("/build")
	g.GET("/pagination", controller.TestPagination)
}
