package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sajan29/data-ingestion/internal/models"
)

func UploadToS3(data []models.Post) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
		Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
		S3ForcePathStyle: aws.Bool(true),
	}))

	svc := s3.New(sess)

	jsonData, _ := json.MarshalIndent(data, "", "  ")

	key := fmt.Sprintf("ingestion-%d.json", data[0].ID)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String(key),
		Body:   bytes.NewReader(jsonData),
	})

	return err
}