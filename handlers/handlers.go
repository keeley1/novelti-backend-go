package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/models"
	"github.com/keeley1/novelti-backend-go/utils"
)

// GetTestSearchHandler function tests calling the google books api.
// It will eventually be deleted.
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
	err = utils.CloseBody(resp.Body)
	if err != nil {
		log.Printf("warning: failed to close response body: %v\n", err)
	}

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
