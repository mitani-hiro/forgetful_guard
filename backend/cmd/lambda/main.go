package main

import (
	"context"
	"encoding/json"
	"log"

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
	err := json.Unmarshal(event.Detail, &geofenceEvent)
	if err != nil {
		log.Printf("JSONパースエラー: %v", err)
		return
	}

	// Geofence イベントの処理
	log.Printf("デバイス %s が %s に %s\n", geofenceEvent.DeviceId, geofenceEvent.GeofenceId, geofenceEvent.EventType)
	log.Printf("Timestamp: %s\n", geofenceEvent.Timestamp)
}

func main() {
	lambda.Start(handler)
}
