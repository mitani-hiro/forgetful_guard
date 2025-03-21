package oapi

// TODO 仮定義（open-apiに移行予定）
type TrackerRequest struct {
	DeviceToken string    `json:"deviceToken"`
	Position    []float64 `json:"position"`
}
