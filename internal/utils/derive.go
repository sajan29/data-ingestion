package utils

import (
	"net/url"
	"strings"
)

func DeriveSourceName(endpoint string) (string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Hostname(), ".")
	if len(parts) >= 1 {
		return parts[0], nil
	}
	return "unknown_source", nil
}