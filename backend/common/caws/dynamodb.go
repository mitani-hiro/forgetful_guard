package caws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

var DynamoDBClient *dynamodb.DynamoDB

// DynamoDB クライアントを初期化して返す
func NewDynamoDBClient() error {
	awsConfig := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}

	if endpoint := os.Getenv("DYNAMODB_ENDPOINT"); endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)
		awsConfig.Credentials = credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return errors.Wrap(err, "session.NewSession error")
	}

	DynamoDBClient = dynamodb.New(sess)
	return nil
}
