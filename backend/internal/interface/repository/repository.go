package repository

import (
	"context"
	"forgetful-guard/internal/domain"
)

// GeofenceRepository インターフェース
type GeofenceRepository interface {
	PutDeviceToken(userID uint64, token string) error
	PutGeofence(ctx context.Context, geofence *domain.Geofence) error
}

type Repository struct {
	Repository GeofenceRepository
}
