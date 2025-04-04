package usecase

import (
	"context"
	"database/sql"
	"forgetful-guard/internal/domain/models"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/interface/repository"
)

// UsecaseService インターフェース
type UsecaseService interface {
	CreateTask(ctx context.Context, tx *sql.Tx, task *models.Task) error
	CreateGeofence(ctx context.Context, req *oapi.Geofence) error
}

type Usecase struct {
	UsecaseService UsecaseService
	Repository     repository.GeofenceRepository
}

func NewUsecase(service UsecaseService, repo repository.GeofenceRepository) *Usecase {
	return &Usecase{UsecaseService: service, Repository: repo}
}
