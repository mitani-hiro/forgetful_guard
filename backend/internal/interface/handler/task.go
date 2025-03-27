package handler

import (
	"forgetful-guard/common/convert"
	"forgetful-guard/internal/interface/oapi"
	task "forgetful-guard/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := task.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	res := make([]*oapi.Task, len(tasks))
	for i, t := range tasks {
		res[i] = &oapi.Task{
			Id:          t.ID,
			Title:       t.Title,
			Description: convert.ToPointer(t.Description.String),
			Completed:   t.Completed,
		}
	}
	c.JSON(http.StatusOK, res)
}

func (h *TaskHandler) GetTask(c *gin.Context, id string) {
	taskID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	task, err := task.GetTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	res := &oapi.Task{
		Id:          uint64(task.ID),
		Title:       task.Title,
		Description: convert.ToPointer(task.Description.String),
		Completed:   task.Completed,
	}
	c.JSON(http.StatusOK, res)
}
