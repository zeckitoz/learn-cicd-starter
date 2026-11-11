package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		wantKey string
		wantErr string // empty means no error is expected
	}{
		"valid api key": {
			headers: http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			wantKey: "my-secret-key",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tc.headers)

			gotErr := ""
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tc.wantErr {
				t.Fatalf("error mismatch: want %q, got %q", tc.wantErr, gotErr)
			}
			if gotKey != tc.wantKey {
				t.Fatalf("key mismatch: want %q, got %q", tc.wantKey, gotKey)
			}
		})
	}
}
