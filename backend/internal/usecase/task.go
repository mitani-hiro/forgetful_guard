package usecase

import (
	"context"
	"database/sql"
	"forgetful-guard/internal/domain"
	"forgetful-guard/internal/domain/models"
	"forgetful-guard/internal/interface/repository"
)

func GetTasks() ([]*models.Task, error) {
	rep := repository.NewTaskRepository()
	return rep.Get()
}

func GetTask(id uint64) (*models.Task, error) {
	rep := repository.NewTaskRepository()
	return rep.GetByID(id)
}

// CreateTask タスク登録処理.
func (u *Usecase) CreateTask(ctx context.Context, tx *sql.Tx, task *models.Task) error {
	if err := domain.ValidateTask(task); err != nil {
		return err
	}

	rep := repository.NewTaskRepository()
	return rep.Create(ctx, tx, task)
}
