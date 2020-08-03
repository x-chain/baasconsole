package gintool

import (
	"github.com/gin-gonic/gin"
	"github.com/x-chain/baasconsole/baas-core/common/log"
	"time"
)

func Logger() gin.HandlerFunc {
	logClient := log.GetLogger("ginlog", log.INFO)

	return func(c *gin.Context) {
		//开始时间
		start := time.Now()
		//处理请求
		c.Next()
		//结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		//path
		path := c.Request.URL.Path
		//ip
		clientIP := c.ClientIP()
		//方法
		method := c.Request.Method
		//状态
		statusCode := c.Writer.Status()
		logClient.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
