package collector

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/sajan29/data-ingestion/internal/models"
	"github.com/sajan29/data-ingestion/internal/utils"
)

func FetchPosts(apiURL string) ([]models.Post, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("non-200 response from API")
	}

	body, err := io.ReadAll(resp.Body)
	utils.Info(string(body))
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	if err := json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}