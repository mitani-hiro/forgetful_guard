package domain

import (
	"github.com/pkg/errors"
)

type Geofence struct {
	TastID  uint64
	Polygon [][][]float64
}

// ValidateGeofence ジオフェンスのバリデーション.
func ValidateGeofence(geofence *Geofence) error {
	if geofence == nil {
		return errors.New("geofence is nil")
	}

	polygon := geofence.Polygon
	if len(polygon) == 0 {
		return errors.New("geofence polygon length is zero")
	}

	if len(polygon[0]) == 0 {
		return errors.New("geofence polygon length is zero")
	}

	return nil
}
