package router

import (
	"forgetful-guard/internal/interface/handler"
	"forgetful-guard/internal/interface/oapi"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(handler.Interceptor())

	h := handler.NewTaskHandler()
	oapi.RegisterHandlersWithOptions(r, h, oapi.GinServerOptions{
		BaseURL: "/api",
	})

	return r
}
