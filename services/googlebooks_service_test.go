package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestConstructAPIURL(t *testing.T) {
	tests := []struct {
		name              string
		query             string
		searchType        string
		startIndex        int
		expectedSubstring string
	}{
		{
			name:              "Search books query",
			query:             "golang",
			searchType:        "search",
			startIndex:        0,
			expectedSubstring: "q=golang",
		},
		{
			name:              "Default query",
			query:             "fiction",
			searchType:        "genre",
			startIndex:        20,
			expectedSubstring: "q=fiction+books",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := ConstructAPIURL(tt.query, tt.searchType, tt.startIndex)
			if !strings.Contains(url, tt.expectedSubstring) {
				t.Errorf("expected %s to contain %s", url, tt.expectedSubstring)
			}
		})
	}
}

func TestMakeAPICall(t *testing.T) {
	// Set up a mock/test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"items": []}`))
	}))
	defer ts.Close()

	// Set up a mock url builder function
	mockURLBuilder := func(query string, searchType string, startIndex int) string {
		return ts.URL
	}

	mockClient := http.DefaultClient
	resp, err := MakeAPICall(mockURLBuilder, mockClient, "golang", "search", 0)
	if err != nil {
		t.Errorf("make api call failed: %v", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 status code but got %d", resp.StatusCode)
	}
}
