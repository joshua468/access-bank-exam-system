// internal/middleware/logging.go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logrus.WithFields(logrus.Fields{
			"status":    status,
			"latency":   latency,
			"client_ip": c.ClientIP(),
			"method":    c.Request.Method,
			"path":      path,
		}).Info("Request received")
	}
}
