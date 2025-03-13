package handler

import (
	"github.com/gin-gonic/gin"
)

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// sessionID, err := c.Cookie("session_id")
		// if err != nil || sessionID == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "session_id is required"})
		// 	c.Abort()
		// 	return
		// }

		// if !isValidSession(sessionID) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}

// isValidSession
// TODO DynamoDBに存在するか.
func isValidSession(sessionID string) bool {
	return sessionID == "valid_session"
}
