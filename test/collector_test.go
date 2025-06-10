package test

import (
	"testing"

	"github.com/sajan29/data-ingestion/internal/collector"
)

func TestFetchPosts(t *testing.T) {
	posts, err := collector.FetchPosts()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(posts) == 0 {
		t.Error("Expected non-zero number of posts")
	}
}