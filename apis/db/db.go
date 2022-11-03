package db

import (
	"apis/models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var provider *dynamo.DB

const (
	localDbURL    = "http://0.0.0.0:8000"
	defaultRegion = "us-east-1"
)

func InitDb() {
	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if region == "" {
		region = defaultRegion
	}

	config := aws.Config{
		MaxRetries:                    aws.Int(3),
		CredentialsChainVerboseErrors: aws.Bool(true), // for full error logs
		Region:                        aws.String(region),
	}

	if accessKey != "" && accessKeySecret != "" {
		config.Credentials = credentials.NewStaticCredentials(accessKey, accessKeySecret, "")
	} else {
		// static config in case of testing or local-setup
		config.Credentials = credentials.NewStaticCredentials("key", "key", "")
		config.Endpoint = aws.String(localDbURL)
	}

	session := session.Must(session.NewSession(&config))
	provider = dynamo.New(session)
	provider.CreateTable(todoTableName, models.Todo{}).Wait()
}
