package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sajan29/data-ingestion/internal/collector"
	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFetchPosts_Success(t *testing.T) {
	mockPosts := []models.Post{
		{ID: 1, UserID: 1, Title: "Test", Body: "Body"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockPosts)
	}))
	defer server.Close()

	posts, err := collector.FetchPosts(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, mockPosts, posts)
}

func TestFetchPosts_Non200(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something went wrong", http.StatusBadRequest)
	}))
	defer server.Close()

	_, err := collector.FetchPosts(server.URL)

	assert.Error(t, err)
	assert.Equal(t, "non-200 response from API", err.Error())
}

func TestFetchPosts_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not-json"))
	}))
	defer server.Close()

	_, err := collector.FetchPosts(server.URL)

	assert.Error(t, err)
}

func TestFetchPosts_RequestFailure(t *testing.T) {
	// Use a non-routable address to simulate request failure
	_, err := collector.FetchPosts("http://127.0.0.1:0")

	assert.Error(t, err)
}