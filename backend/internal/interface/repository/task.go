package repository

import (
	"context"
	"database/sql"
	"forgetful-guard/internal/domain/models"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepositoryImpl struct{}

func NewTaskRepository() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{}
}

func (r *TaskRepositoryImpl) Get() ([]*models.Task, error) {
	t := []*models.Task{
		{
			ID:          1,
			Title:       "hoge title 1",
			Description: null.StringFrom("hoge hoge 1"),
			Completed:   true,
		},
		{
			ID:          2,
			Title:       "hoge title 2",
			Description: null.StringFrom("hoge hoge 2"),
			Completed:   true,
		},
	}
	return t, nil
}

func (r *TaskRepositoryImpl) GetByID(id uint64) (*models.Task, error) {
	t := &models.Task{
		ID:          id,
		Title:       "hoge title",
		Description: null.StringFrom("hoge hoge"),
		Completed:   true,
	}
	return t, nil
}

// Create タスク登録.
func (r *TaskRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, task *models.Task) error {
	if err := task.Insert(ctx, tx, boil.Infer()); err != nil {
		return errors.Wrap(err, "task insert error")
	}

	return nil
}

func (r *TaskRepositoryImpl) Update(task *models.Task) error {
	return nil
}

func (r *TaskRepositoryImpl) Delete(id uint64) error {
	return nil
}
