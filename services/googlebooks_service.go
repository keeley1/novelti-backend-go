package services

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	"github.com/keeley1/novelti-backend-go/config"
	"github.com/keeley1/novelti-backend-go/models"
)

// URLBuilder is a function type for constructing URL.
// Can be used here and for testing.
type URLBuilder func(query string, searchType string, startIndex int) string

// ConstructAPIURL takes a query, search type, and a start index and constructs
// an API url.
func ConstructAPIURL(query string, searchType string, startIndex int) string {
	var apiURL string
	encodedQuery := ""

	if searchType == string(models.SearchBooks) {
		encodedQuery = url.QueryEscape(query)
	} else {
		encodedQuery = url.QueryEscape(query + " books")
		fmt.Println("encoded query: ", encodedQuery)
	}

	apiURL = fmt.Sprintf("%s?q=%s&orderBy=%s&maxResults=%d&startIndex=%d",
		config.GoogleBooksBaseURL,
		encodedQuery,
		config.DefaultOrderBy,
		config.DefaultMaxResults,
		startIndex)

	return apiURL
}

// MakeAPICall takes a URLBuilder, http client, query, search type, and start
// index and makes a call to the google books API.
func MakeAPICall(buildURL URLBuilder,
	client *http.Client,
	query string,
	searchType string,
	startIndex int) (*http.Response, error) {

	var googleBooksAPIURL = buildURL(query, searchType, startIndex)

	fmt.Println("Calling google books api.....")
	resp, err := client.Get(googleBooksAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}

	return resp, err
}

// DecodeBookData takes in a http response and deserializes it
// into a response object.
func DecodeBookData(responseData *http.Response) (*models.GoogleBooksAPIResponse, error) {
	// Read response body into byte slice
	body, err := io.ReadAll(responseData.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var apiResp models.GoogleBooksAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &apiResp, nil
}

// CreateBooks takes the API response object and creates book objects.
func CreateBooks(decodedBooks *models.GoogleBooksAPIResponse) ([]models.Book, error) {
	var books []models.Book
	for _, item := range decodedBooks.Items {
		authors := item.VolumeInfo.Authors
		if len(authors) == 0 {
			authors = []string{"Author Unknown"}
		}

		book := models.Book{
			ID:            item.ID,
			Title:         item.VolumeInfo.Title,
			Authors:       authors,
			PublishedDate: item.VolumeInfo.PublishedDate,
			Cover:         item.VolumeInfo.ImageLinks.Thumbnail,
			Description:   template.HTML(item.VolumeInfo.Description),
		}
		books = append(books, book)
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("no valid books found")
	}

	return books, nil
}
