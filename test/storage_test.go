package test

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/sajan29/data-ingestion/internal/storage"
	"github.com/stretchr/testify/assert"
)

// Mock S3 client
type mockS3Client struct {
	s3iface.S3API
	PutObjectFunc func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

func (m *mockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return m.PutObjectFunc(input)
}

func TestUploadToS3_Success(t *testing.T) {
	os.Setenv("S3_BUCKET_NAME", "test-bucket")

	mockClient := &mockS3Client{
		PutObjectFunc: func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(input.Body)
			assert.Contains(t, buf.String(), `"id": 1`)
			assert.Equal(t, "test-bucket", *input.Bucket)
			return &s3.PutObjectOutput{}, nil
		},
	}

	sampleData := []models.Post{{ID: 1, Title: "test", Body: "sample body", UserID: 123}}
	err := storage.UploadToS3(mockClient, sampleData)
	assert.NoError(t, err)
}

func TestUploadToS3_Failure(t *testing.T) {
	os.Setenv("S3_BUCKET_NAME", "test-bucket")

	mockClient := &mockS3Client{
		PutObjectFunc: func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			return nil, errors.New("S3 failure")
		},
	}

	sampleData := []models.Post{{ID: 2, Title: "fail", Body: "bad", UserID: 999}}
	err := storage.UploadToS3(mockClient, sampleData)
	assert.Error(t, err)
	assert.Equal(t, "S3 failure", err.Error())
}