package usecase

import (
	"context"
	"fmt"
	"forgetful-guard/internal/domain"
	"forgetful-guard/internal/interface/oapi"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pkg/errors"
)

// SendTracker
func SendTracker(ctx context.Context, req *oapi.TrackerRequest) error {
	tracker := &domain.Tracker{
		DeviceToken: req.DeviceToken,
		Position:    req.Position,
	}

	if err := domain.ValidateTracker(tracker); err != nil {
		return err
	}

	if err := sendTracker(ctx, tracker); err != nil {
		return err
	}

	return nil
}

// sendTracker トラッカー送信.
func sendTracker(ctx context.Context, tracker *domain.Tracker) error {
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

	res, err := svc.BatchUpdateDevicePosition(ctx, &location.BatchUpdateDevicePositionInput{
		TrackerName: aws.String("ForgetfulGuardTracker"),
		Updates: []types.DevicePositionUpdate{
			{
				DeviceId:   aws.String(tracker.DeviceToken),
				Position:   tracker.Position,
				SampleTime: aws.Time(time.Now()),
			},
		},
	})
	if err != nil {
		return errors.Wrap(err, "BatchUpdateDevicePosition error")
	}

	fmt.Printf("sendTracker チェック完了: %+v\n", res.Errors)
	return nil
}
