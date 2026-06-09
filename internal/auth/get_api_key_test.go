package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	t.Run("returns error when no auth header is present", func(t *testing.T) {
		headers := make(http.Header)
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})

	t.Run("returns error when auth header is present but not in the correct format", func(t *testing.T) {
		headers := make(http.Header)
		headers.Add("Authorization", "Bearer 1234567890")
		_, err := GetAPIKey(headers)
		if err.Error() != "malformed authorization header" {
			t.Errorf("expected error %v, got %v", "malformed authorization header", err)
		}
	})
	t.Run("returns api key when auth header is present and in the correct format", func(t *testing.T) {
		headers := make(http.Header)
		headers.Add("Authorization", "ApiKey 1234567890")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if apiKey != "1234567890" {
			t.Errorf("expected api key to be 1234567890, got %s", apiKey)
		}
	})
}