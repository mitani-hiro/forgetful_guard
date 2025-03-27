package interceptor

import (
	"forgetful-guard/common/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery リカバリーインターセプター
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Recovered from panic", err)

				c.JSON(http.StatusInternalServerError, gin.H{"error": "内部サーバーエラーが発生しました"})

				c.Abort()
			}
		}()

		c.Next()
	}
}
