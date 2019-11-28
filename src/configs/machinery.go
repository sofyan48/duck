package configs

import (
	"net/http"
	"os"
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
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET_KEY"), ""),
		HTTPClient: &http.Client{
			Timeout: time.Second * 120,
		},
	})))

	var visibilityTimeout = 20
	cnf := &config.Config{
		Broker:        os.Getenv("SQS_BROKER"),
		DefaultQueue:  os.Getenv("SQS_SCHEDULER"),
		ResultBackend: os.Getenv("RESULT_BACKEND"),
		SQS: &config.SQSConfig{
			Client:            sqsClient,
			VisibilityTimeout: &visibilityTimeout,
			WaitTimeSeconds:   30,
		},
	}
	return cnf, nil
}
