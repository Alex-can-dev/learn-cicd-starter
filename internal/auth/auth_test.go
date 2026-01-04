package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		header        string
		expectedKey   string
		expectedError error
	}{
		{"missing header", "", "", ErrNoAuthHeaderIncluded},
		{"wrong format", "Bearer 12345", "", errors.New("malformed authorization header")},
		{"no key", "ApiKey", "", errors.New("malformed authorization header")},
		{"valid header", "ApiKey 12345", "12345", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}
			key, err := GetAPIKey(headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %s, got %s", tt.expectedKey, key)
			}
			if (err == nil && tt.expectedError != nil) || (err != nil && tt.expectedError == nil) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			} else if err != nil && tt.expectedError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
