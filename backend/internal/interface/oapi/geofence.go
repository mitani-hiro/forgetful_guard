package oapi

// TODO 仮定義（open-apiに移行予定）
type GeofenceRequest struct {
	Title       string        `json:"title"`
	UserID      uint64        `json:"userID"`
	Polygon     [][][]float64 `json:"polygon"`
	DeviceToken string        `json:"deviceToken"`
}
