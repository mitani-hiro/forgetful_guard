package repository

import (
	"context"
	"forgetful-guard/common/caws"
	"forgetful-guard/common/logger"
	"forgetful-guard/internal/domain"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

// PutDeviceToken デバイス情報登録.
func (r *Repository) PutDeviceToken(userID uint64, token string) error {
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

// PutGeofence ジオフェンス登録.
func (r *Repository) PutGeofence(ctx context.Context, geofence *domain.Geofence) error {
	geofenceEntry := types.BatchPutGeofenceRequestEntry{
		GeofenceId: aws.String(strconv.FormatUint(geofence.TastID, 10)),
		Geometry: &types.GeofenceGeometry{
			Polygon: geofence.Polygon,
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

	res, err := caws.LocationClient.BatchPutGeofence(ctx, input)
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
