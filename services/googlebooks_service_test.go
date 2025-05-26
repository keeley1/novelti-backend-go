package services

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/keeley1/novelti-backend-go/utils"
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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"items": []}`))
	}))
	defer ts.Close()

	// Set up a mock url builder function
	mockURLBuilder := func(_ string, _ string, _ int) string {
		return ts.URL
	}

	mockClient := http.DefaultClient
	resp, err := MakeAPICall(mockURLBuilder, mockClient, "golang", "search", 0)
	if err != nil {
		t.Errorf("make api call failed: %v", err)
	}
	err = utils.CloseBody(resp.Body)
	if err != nil {
		t.Errorf("warning: failed to close response body: %v\n", err)
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 status code but got %d", resp.StatusCode)
	}
}

func TestDecodeBookData_Success(t *testing.T) {
	// JSON matching GoogleBooksAPIResponse structure
	jsonData := `{
		"items": [
			{
				"id": "123",
				"volumeInfo": {
					"title": "Sample Book",
					"authors": ["Author One"],
					"publishedDate": "2020-01-01",
					"description": "A sample book description.",
					"imageLinks": {
						"thumbnail": "http://example.com/thumb.jpg"
					}
				}
			}
		]
	}`

	// Create a fake http.Response with the JSON in the Body
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(jsonData)),
	}

	result, err := DecodeBookData(resp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(result.Items))
	}
}

func TestDecodeBookData_WithInvalidJSON(t *testing.T) {
	invalidJSON := `{"kind": "books#volumes", "totalItems": 1, "items": [INVALID]}`

	// Create a fake http.Response with the JSON in the Body
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(invalidJSON)),
	}

	_, err := DecodeBookData(resp)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}
