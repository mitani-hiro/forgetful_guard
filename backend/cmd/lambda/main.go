package main

import (
	"context"
	"encoding/json"
	"forgetful-guard/common/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type GeofenceEvent struct {
	EventType  string `json:"EventType"`
	GeofenceId string `json:"GeofenceId"`
	DeviceId   string `json:"DeviceId"`
	Timestamp  string `json:"SampleTime"`
}

func handler(ctx context.Context, event events.EventBridgeEvent) {
	var geofenceEvent GeofenceEvent
	if err := json.Unmarshal(event.Detail, &geofenceEvent); err != nil {
		logger.Error("json.Unmarshal error", err)
		return
	}

	// Geofence イベントの処理
	logger.Info("tracker event received", "deviceID", geofenceEvent.DeviceId, "geofenceID", geofenceEvent.GeofenceId, "eventType", geofenceEvent.EventType, "timestamp", geofenceEvent.Timestamp)
}

func main() {
	lambda.Start(handler)
}
