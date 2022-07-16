package middlewares

import (
	uuid "github.com/satori/go.uuid"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddTraceId() gin.HandlerFunc {
	return func(g *gin.Context) {
		traceId := g.GetHeader("traceId")
		if traceId == "" {
			traceId = uuid.NewV4().String()
			g.Set("traceId", traceId)
		}
		g.Request = g.Request.WithContext(g)
		//log.Info("AddTraceId success")
		g.Next()
	}
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 请求路径
		path := c.Request.URL.Path
		// 请求参数
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		// 若response的状态码不是200为异常
		if c.Writer.Status() != 200 {
			// 记录异常信息
			zap.L().Info(path,
				zap.String("traceId", c.GetString("traceId")),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}
	}
}
