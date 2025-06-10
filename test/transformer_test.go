package test

import (
	"testing"
	"time"

	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/sajan29/data-ingestion/internal/transformer"
)

func TestTransform(t *testing.T) {
	posts := []models.Post{{ID: 1, Title: "Test"}}
	source := "test_source"
	result := transformer.Transform(posts, source)

	if result[0].Source != source {
		t.Error("Source not set correctly")
	}

	_, err := time.Parse(time.RFC3339, result[0].IngestedAt)
	if err != nil {
		t.Error("Invalid timestamp format")
	}
}