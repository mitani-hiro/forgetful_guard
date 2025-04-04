package domain

import (
	"forgetful-guard/internal/domain/models"
	"unicode/utf8"

	"github.com/pkg/errors"
)

// ValidateTask タスクのバリデーション.
func ValidateTask(task *models.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	if task.Title == "" {
		return errors.Errorf("task title is empty, task_id: %v", task.ID)
	}

	if utf8.RuneCountInString(task.Title) >= 100 {
		return errors.Errorf("task title is more than 100, task_id: %v", task.ID)
	}

	if task.UserID == 0 {
		return errors.Errorf("task user id is zero, task_id: %v", task.ID)
	}

	return nil
}
