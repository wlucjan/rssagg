package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API key from header of a HTTP request
// Example: Authorization: ApiKey <api_key>
func GetAPIKey(r *http.Request) (string, error) {
	apiKey := r.Header.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no authentication info found")
	}

	if !strings.HasPrefix(apiKey, "ApiKey ") {
		return "", errors.New("malformed auth header")
	}

	apiKey = strings.TrimPrefix(apiKey, "ApiKey ")

	return apiKey, nil
}
