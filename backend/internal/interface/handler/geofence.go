package handler

import (
	"forgetful-guard/common/logger"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostGeofence ジオフェンス登録.
func (h *TaskHandler) PostGeofence(c *gin.Context) {
	// TODO 認証処理

	var req *oapi.Geofence
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("c.ShouldBindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "c.ShouldBindJSON error"})
		return
	}

	if err := usecase.CreateGeofence(c, req); err != nil {
		logger.Error("usecase.CreateGeofence error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase.CreateGeofence error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
