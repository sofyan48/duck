package configs

import (
	"net/http"
	"time"

	"github.com/RichardKnop/machinery/v1/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// LoadMachineryConfig load config Function
func LoadMachineryConfig(path string) (*config.Config, error) {
	var sqsClient = sqs.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIAJT6VRYLHGO2NMU7Q", "J8ndXGNUr+FyqTnWJAmYjT3G+OaGwkAkdpgb3f6m", ""),
		HTTPClient: &http.Client{
			Timeout: time.Second * 120,
		},
	})))

	var visibilityTimeout = 20
	cnf := &config.Config{
		Broker:        "https://sqs.ap-southeast-1.amazonaws.com/025432561883",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "redis://localhost:6379",
		SQS: &config.SQSConfig{
			Client:            sqsClient,
			VisibilityTimeout: &visibilityTimeout,
			WaitTimeSeconds:   30,
		},
	}
	return cnf, nil
}
