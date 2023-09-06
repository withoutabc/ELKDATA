package middleware

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gin-gonic/gin"
	"time"
)

func ReqTimeInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("in")
		start := time.Now()
		c.Next()
		cost := time.Since(start)
		hlog.Infof("cost:%v", cost)
	}
}
