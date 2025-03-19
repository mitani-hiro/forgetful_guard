package domain

import (
	"github.com/pkg/errors"
)

type Tracker struct {
	DeviceID string
	Position []float64
}

// ValidateTracker トラッカーのバリデーション.
func ValidateTracker(tracker *Tracker) error {
	if tracker == nil {
		return errors.New("tracker is nil")
	}

	if tracker.DeviceID == "" {
		return errors.New("tracker deviceID is empty")
	}

	if len(tracker.Position) != 2 {
		return errors.New("tracker position is invalid")
	}

	return nil
}
