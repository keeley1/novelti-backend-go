package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/models"
)

// check queries with spaces

func GetTestSearchHandler(context *gin.Context) {
	searchQuery := context.Param("query")
	googleBooksAPIURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s&maxResults=20", searchQuery)

	resp, err := http.Get(googleBooksAPIURL)
	if err != nil {
		// Respond 500 - internal server error
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to fetch data from Google Books API"},
		)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// Respond 500 - internal server error
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to parse response body"},
		)
		return
	}

	fmt.Println(googleBooksAPIURL)
	context.JSON(http.StatusOK, data)
}

func GetBooksByTitleHandler(context *gin.Context) {
	searchQuery := context.Param("title")
	searchType := models.SearchByTitle
	HandleBookSearch(context, searchQuery, string(searchType))
}

func GetBooksByGenreHandler(context *gin.Context) {
	searchQuery := context.Param("genre")
	searchType := models.SearchByGenre
	HandleBookSearch(context, searchQuery, string(searchType))
}

func GetBooksBySearchHandler(context *gin.Context) {
	searchQuery := context.Param("searchquery")
	searchType := models.SearchBooks
	HandleBookSearch(context, searchQuery, string(searchType))
}
