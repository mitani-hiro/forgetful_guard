package caws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/location"
	"github.com/pkg/errors"
)

var LocationClient *location.Client

// NewLocationClient 実際の AWS Location Service クライアントを作成.
func NewLocationClient() error {
	optFns := []func(*config.LoadOptions) error{
		config.WithRegion(os.Getenv("AWS_REGION")),
	}

	if os.Getenv("APP_ENV") == "local" {
		optFns = append(optFns, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")))
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), optFns...)
	if err != nil {
		return errors.Wrap(err, "config.LoadDefaultConfig error")
	}

	LocationClient = location.NewFromConfig(cfg)
	return nil
}
