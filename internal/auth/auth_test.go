package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct{
		name		string
		headerValue	string
		expectAPIKey string 
		expectError  bool
	}{
		{
			name:	"valid header",
			headerValue: "ApiKey abcd1234",
			expectAPIKey: "abcd1234",
			expectError: false,
		},
		{
			name: "missing header",
			headerValue: "",
			expectError: true,
		},
		{
			name: "wrong header",
			headerValue: "Bearer abcdefg1234",
			expectError: true,
		},
		{
			name: "missing api key",
			headerValue: "ApiKey",
			expectError: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			headers := http.Header{}

			if testCase.headerValue != "" {
				headers.Set("Authorization", testCase.headerValue)
			}

			apiKey, err := GetAPIKey(headers)
			if testCase.expectError {
				if err == nil {
					t.Error("expect error but got nil")
		
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpect error: %v", err)
			}

			if apiKey != testCase.expectAPIKey {
				t.Errorf("expect apikey: %s, got: %s", testCase.expectAPIKey, apiKey)
			}
		
		})
	}
}