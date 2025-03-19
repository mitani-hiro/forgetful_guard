package oapi

// TODO 仮定義（open-apiに移行予定）
type TrackerRequest struct {
	DeviceID string    `json:"deviceID"`
	Position []float64 `json:"position"`
}
