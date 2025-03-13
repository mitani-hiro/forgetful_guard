package handler

import (
	"context"
	"fmt"
	"forgetful-guard/internal/interface/oapi"
	"forgetful-guard/internal/usecase"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/gin-gonic/gin"
)

func CheckGeofenceSample(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return
	}

	svc := location.NewFromConfig(cfg)

	res, err := svc.BatchUpdateDevicePosition(context.TODO(), &location.BatchUpdateDevicePositionInput{
		TrackerName: aws.String("explore.tracker"),
		Updates: []types.DevicePositionUpdate{
			{
				DeviceId:   aws.String("test-device"),
				Position:   []float64{139.6950, 35.6925},
				SampleTime: aws.Time(time.Now()),
			},
		},
	})
	if err != nil {
		log.Fatalf("Geofence 判定エラー: %v", err)
		return
	}

	fmt.Printf("Geofence チェック完了: %+v\n", res.Errors)
}

func CreateGeofenceSample(c *gin.Context) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return
	}

	svc := location.NewFromConfig(cfg)

	polygon := [][][]float64{
		{
			{139.6917, 35.6895}, // 東京の座標
			{139.7000, 35.6895},
			{139.7000, 35.6950},
			{139.6917, 35.6950},
			{139.6917, 35.6895}, // 最初の座標と一致させる
		},
	}

	// 逆順に並べる
	reversedPolygon := ensureCounterClockwise(polygon)

	geofenceID := "task-102" // タスクID
	geofenceEntry := types.BatchPutGeofenceRequestEntry{
		GeofenceId: aws.String(geofenceID),
		Geometry: &types.GeofenceGeometry{
			Polygon: reversedPolygon,
		},
		GeofenceProperties: map[string]string{
			"task_name": "Sample Task",
			"priority":  "high",
		},
	}

	input := &location.BatchPutGeofenceInput{
		CollectionName: aws.String("explore.geofence-collection"),
		Entries:        []types.BatchPutGeofenceRequestEntry{geofenceEntry},
	}

	res, err := svc.BatchPutGeofence(context.TODO(), input)
	if err != nil {
		log.Fatalf("Geofence 登録エラー: %v", err)
	}

	if len(res.Errors) > 0 {
		for _, e := range res.Errors {
			log.Printf("エラー: GeofenceId=%s, メッセージ=%s", *e.GeofenceId, *e.Error.Message)
		}
	}

	fmt.Printf("Geofence ResultMetadata: %+v\n", res.ResultMetadata)
	fmt.Printf("Geofence Successes: %+v\n", res.Successes)
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

// CreateGeofence ジオフェンス登録.
func CreateGeofence(c *gin.Context) {
	// TODO 認証処理

	var req *oapi.GeofenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("c.ShouldBindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "c.ShouldBindJSON error"})
		return
	}

	if err := usecase.CreateGeofence(c, req); err != nil {
		fmt.Printf("usecase.CreateGeofence error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "usecase.CreateGeofence error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
