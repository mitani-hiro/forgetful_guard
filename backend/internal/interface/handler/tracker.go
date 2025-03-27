package handler

import (
	"forgetful-guard/common/logger"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostTracker トラッカー送信.
func (h *TaskHandler) PostTracker(c *gin.Context) {
	// TODO 認証処理

	var req *oapi.Tracker
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("c.ShouldBindJSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "c.ShouldBindJSON error"})
		return
	}

	if err := usecase.SendTracker(c, req); err != nil {
		logger.Error("usecase.SendTracker error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase.SendTracker error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
