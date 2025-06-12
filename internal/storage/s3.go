package storage

import (
	"time"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/sajan29/data-ingestion/internal/models"
)

// CreateS3Client returns a real S3 client (implements s3iface.S3API)
func CreateS3Client() s3iface.S3API {

	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	}

	if endpoint := os.Getenv("S3_ENDPOINT"); endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
		cfg.S3ForcePathStyle = aws.Bool(true) // required for LocalStack
	}

	sess := session.Must(session.NewSession(cfg))
	return s3.New(sess)
}

// UploadToS3 uploads JSON data to an S3 bucket
func UploadToS3(svc s3iface.S3API, data []models.Post) error {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	timestamp := time.Now().UTC().Format("20060102T150405")

	key := fmt.Sprintf("ingestion-%s.json", timestamp)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
		Body:   bytes.NewReader(jsonData),
	})

	return err
}