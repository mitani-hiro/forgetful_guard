package handler

import (
	task "forgetful-guard/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func GetTasks(c *gin.Context) {
	tasks, err := task.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	res := make([]*Task, len(tasks))
	for i, t := range tasks {
		res[i] = &Task{
			ID:          uint64(t.ID),
			Title:       t.Title,
			Description: t.Description.String,
			Completed:   t.Completed,
		}
	}
	c.JSON(http.StatusOK, res)
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
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

	res := &Task{
		ID:          uint64(task.ID),
		Title:       task.Title,
		Description: task.Description.String,
		Completed:   task.Completed,
	}
	c.JSON(http.StatusOK, res)
}

func CreateTask(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}

func UpdatesTask(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}

func DeleteTask(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{})
}
