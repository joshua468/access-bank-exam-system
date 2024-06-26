// internal/middleware/error.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/access-bank-exam-system/pkg/utils"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			appErr := utils.HandleError(err)
			c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		}
	}
}
