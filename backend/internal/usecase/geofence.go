package usecase

import (
	"context"
	"database/sql"
	"forgetful-guard/common/caws"
	"forgetful-guard/common/logger"
	"forgetful-guard/common/rdb"
	"forgetful-guard/internal/domain"
	"forgetful-guard/internal/domain/models"
	"forgetful-guard/internal/interface/oapi"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

// CreateGeofence
func CreateGeofence(ctx context.Context, req *oapi.Geofence) error {
	return rdb.Tx(ctx, func(tx *sql.Tx) error {
		task := &models.Task{
			Title:  req.Title,
			UserID: req.UserID,
		}

		if err := CreateTask(ctx, tx, task); err != nil {
			return err
		}

		geofence := &domain.Geofence{
			TastID:  task.ID,
			Polygon: req.Polygon,
		}

		if err := domain.ValidateGeofence(geofence); err != nil {
			return err
		}

		if err := putGeofence(ctx, geofence); err != nil {
			return err
		}

		if err := putDeviceToken(task.UserID, req.DeviceToken); err != nil {
			return err
		}

		return nil
	})
}

// putDeviceToken デバイス情報登録.
func putDeviceToken(userID uint64, token string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("device_tokens"),
		Item: map[string]*dynamodb.AttributeValue{
			"user_id": {
				N: aws.String(strconv.FormatUint(userID, 10)),
			},
			"device_token": {
				S: aws.String(token),
			},
		},
	}

	if _, err := caws.DynamoDBClient.PutItem(input); err != nil {
		return errors.Wrap(err, "dynamoDB put error")
	}

	return nil
}

// putGeofence ジオフェンス登録.
func putGeofence(ctx context.Context, geofence *domain.Geofence) error {
	optFns := []func(*config.LoadOptions) error{
		config.WithRegion(os.Getenv("AWS_REGION")),
	}

	if os.Getenv("APP_ENV") == "local" {
		optFns = append(optFns, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")))
	}

	cfg, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		return errors.Wrap(err, "config.LoadDefaultConfig error")
	}

	svc := location.NewFromConfig(cfg)

	geofenceEntry := types.BatchPutGeofenceRequestEntry{
		GeofenceId: aws.String(strconv.FormatUint(geofence.TastID, 10)),
		Geometry: &types.GeofenceGeometry{
			Polygon: ensureCounterClockwise(geofence.Polygon),
		},
		GeofenceProperties: map[string]string{
			"task_name": "Sample Task",
			"priority":  "high",
		},
	}

	input := &location.BatchPutGeofenceInput{
		CollectionName: aws.String("ForgetfulGuardGeofenceCollection"),
		Entries:        []types.BatchPutGeofenceRequestEntry{geofenceEntry},
	}

	res, err := svc.BatchPutGeofence(ctx, input)
	if err != nil {
		return errors.Wrap(err, "BatchPutGeofence error")
	}

	if len(res.Errors) > 0 {
		for _, e := range res.Errors {
			logger.Info("response errors", "geofenceID", *e.GeofenceId, "message", *e.Error.Message)
		}
	}

	logger.Info("Geofence Successes", "successes", res.Successes)
	return nil
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
