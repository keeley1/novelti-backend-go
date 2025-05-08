package services

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	"github.com/keeley1/novelti-backend-go/models"
)

// Could make more generic and move to utils??
func constructAPIURL(query string, searchType string, startIndex int) string {
	var googleBooksAPIURL string

	if searchType == string(models.SearchId) {
		googleBooksAPIURL = fmt.Sprintf("https://www.googleapis.com/books/v1/volumes/%s", query)
	} else if searchType == string(models.SearchBooks) {
		encodedQuery := url.QueryEscape(query)
		googleBooksAPIURL = fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&orderBy=newest&maxResults=40&startIndex=%d", encodedQuery, startIndex)
	} else {
		encodedQuery := url.QueryEscape(query + " books")
		googleBooksAPIURL = fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&orderBy=newest&maxResults=40&startIndex=%d", encodedQuery, startIndex)
	}

	fmt.Println("API URL:", googleBooksAPIURL)
	return googleBooksAPIURL
}

func MakeAPICall(query string, searchType string, startIndex int) (*http.Response, error) {
	var googleBooksAPIURL = constructAPIURL(query, searchType, startIndex)

	fmt.Println("Calling google books api.....")
	resp, err := http.Get(googleBooksAPIURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	return resp, err
}

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

	for _, item := range apiResp.Items {
		fmt.Println("Title:", item.VolumeInfo.Title)
		fmt.Println("Authors:", item.VolumeInfo.Authors)
		fmt.Println("Thumbnail:", item.VolumeInfo.ImageLinks.Thumbnail)
	}
	return &apiResp, nil
}

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
