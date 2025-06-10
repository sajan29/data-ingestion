package transformer

import (
	"time"

	"github.com/sajan29/data-ingestion/internal/models"
)

func Transform(posts []models.Post, source string) []models.Post {
	for i := range posts {
		posts[i].IngestedAt = time.Now().UTC().Format(time.RFC3339)
		posts[i].Source = source
	}
	return posts
}