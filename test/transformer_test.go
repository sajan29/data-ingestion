package test

import (
	"testing"
	"time"

	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/sajan29/data-ingestion/internal/transformer"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	source := "jsonplaceholder"

	input := []models.Post{
		{ID: 1, Title: "Test 1", Body: "Body 1", UserID: 100},
		{ID: 2, Title: "Test 2", Body: "Body 2", UserID: 101},
	}

	result := transformer.Transform(input, source)

	// Check length
	assert.Equal(t, len(input), len(result))

	for _, post := range result {
		// Check Source is correctly set
		assert.Equal(t, source, post.Source)

		// Check IngestedAt is a valid RFC3339 timestamp
		_, err := time.Parse(time.RFC3339, post.IngestedAt)
		assert.NoError(t, err, "IngestedAt should be a valid RFC3339 timestamp")
	}
}