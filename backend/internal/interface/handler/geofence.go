package handler

import (
	"fmt"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGeofence ジオフェンス登録.
func CreateGeofence(c *gin.Context) {
	// TODO 認証処理

	var req *oapi.GeofenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("c.ShouldBindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "c.ShouldBindJSON error"})
		return
	}

	if err := usecase.CreateGeofence(c, req); err != nil {
		fmt.Printf("usecase.CreateGeofence error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase.CreateGeofence error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
