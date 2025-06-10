package test

import (
	"os"
	"testing"
	"time"

	"github.com/sajan29/data-ingestion/internal/config"
	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/sajan29/data-ingestion/internal/storage"
)

func TestUploadToS3(t *testing.T) {
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET", "my-test-bucket")

	cfg := config.LoadEnv()

	mockData := []models.Post{
		{
			Post: models.Post{
				UserID: 1,
				ID:     12345,
				Title:  "Hello",
				Body:   "World",
			},
			IngestedAt: time.Now().Format(time.RFC3339),
			Source:     "test",
		},
	}

	err := storage.UploadToS3(cfg, mockData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}