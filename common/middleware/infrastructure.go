package middleware

import (
	"github.com/faiz/go-mall/common/util"
	"github.com/gin-gonic/gin"
)

func StartTrace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("traceId")
		parentId := c.GetHeader("parentId")
		spanId := util.GenerateSpanId(c.Request.RemoteAddr)
		if traceId == "" {
			traceId = spanId
		}
		c.Set("traceId", traceId)
		c.Set("parentId", parentId)
		c.Set("spanId", spanId)
		c.Next()
	}
}
