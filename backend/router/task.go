package router

import (
	"forgetful-guard/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func setTaskRouter(r *gin.Engine) {
	// タスク一覧
	r.GET("/tasks", handler.GetTasks)
	// タスク取得
	r.GET("/task/:id", handler.GetTask)
	// タスク登録
	r.POST("/task", handler.CreateTask)
	// タスク更新
	r.PUT("/task/:id", handler.UpdatesTask)
	// タスク削除
	r.DELETE("/task/:id", handler.DeleteTask)
}
