package caws

import (
	"log"
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
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Endpoint:    aws.String("http://dynamodb-local-example:8000"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
		return errors.Wrap(err, "session.NewSession error")
	}

	DynamoDBClient = dynamodb.New(sess)
	return nil
}
