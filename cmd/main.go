package main

import (
	"log"

	"github.com/sajan29/data-ingestion/internal/collector"
	"github.com/sajan29/data-ingestion/internal/config"
	"github.com/sajan29/data-ingestion/internal/storage"
	"github.com/sajan29/data-ingestion/internal/transformer"
	"github.com/sajan29/data-ingestion/internal/utils"
)

func main() {
	config.LoadEnv()

	apiEndpoint := config.GetEnv("API_ENDPOINT")
	source, err := utils.DeriveSourceName(apiEndpoint)
	if err != nil {
		log.Fatalf("failed to derive source: %v", err)
	}

	posts, err := collector.FetchPosts(apiEndpoint)
	if err != nil {
		log.Fatalf("failed to fetch posts: %v", err)
	}

	transformed := transformer.Transform(posts, source)
    svc := storage.CreateS3Client()
	err = storage.UploadToS3(svc, transformed)
	if err != nil {
		log.Fatalf("upload failed: %v", err)
	}

	log.Println("Ingestion complete.")
}