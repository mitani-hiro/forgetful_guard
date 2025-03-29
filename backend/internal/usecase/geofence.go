package usecase

import (
	"context"
	"database/sql"
	"forgetful-guard/common/rdb"
	"forgetful-guard/internal/domain"
	"forgetful-guard/internal/domain/models"
	"forgetful-guard/internal/interface/oapi"
)

// CreateGeofence
func (u *Usecase) CreateGeofence(ctx context.Context, req *oapi.Geofence) error {
	return rdb.Tx(ctx, func(tx *sql.Tx) error {
		task := &models.Task{
			Title:  req.Title,
			UserID: req.UserID,
		}

		if err := u.UsecaseService.CreateTask(ctx, tx, task); err != nil {
			return err
		}

		geofence := &domain.Geofence{
			TastID:  task.ID,
			Polygon: req.Polygon,
		}

		if err := domain.ValidateGeofence(geofence); err != nil {
			return err
		}

		geofence.Polygon = ensureCounterClockwise(geofence.Polygon)
		if err := u.Repository.PutGeofence(ctx, geofence); err != nil {
			return err
		}

		if err := u.Repository.PutDeviceToken(task.UserID, req.DeviceToken); err != nil {
			return err
		}

		return nil
	})
}

func ensureCounterClockwise(polygon [][][]float64) [][][]float64 {
	for i := range polygon {
		if isClockwise(polygon[i]) {
			// 時計回りなら反転
			for j, k := 0, len(polygon[i])-1; j < k; j, k = j+1, k-1 {
				polygon[i][j], polygon[i][k] = polygon[i][k], polygon[i][j]
			}
		}
	}
	return polygon
}

func isClockwise(polygon [][]float64) bool {
	sum := 0.0
	n := len(polygon)

	for i := 0; i < n-1; i++ {
		sum += (polygon[i+1][0] - polygon[i][0]) * (polygon[i+1][1] + polygon[i][1])
	}
	sum += (polygon[0][0] - polygon[n-1][0]) * (polygon[0][1] + polygon[n-1][1])

	return sum > 0 // 時計回りなら true、反時計回りなら false
}
