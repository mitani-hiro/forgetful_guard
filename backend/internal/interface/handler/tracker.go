package handler

import (
	"fmt"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendTracker(c *gin.Context) {
	// TODO 認証処理

	var req *oapi.TrackerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("c.ShouldBindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "c.ShouldBindJSON error"})
		return
	}

	fmt.Printf("SendTracker req: %+v\n", *req)

	if err := usecase.SendTracker(c, req); err != nil {
		fmt.Printf("usecase.SendTracker error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase.SendTracker error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
