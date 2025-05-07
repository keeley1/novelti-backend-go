package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/models"
	"github.com/keeley1/novelti-backend-go/services"
	"github.com/keeley1/novelti-backend-go/utils"
)

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
	startIndexStr := context.DefaultQuery("startIndex", "0")
	startIndex := utils.ParseToPositiveInt(startIndexStr)

	// Ignore caching for now

	// Make api call
	resp, err := services.MakeAPICall(searchQuery, "intitle", startIndex)
	if err != nil {
		log.Printf("Google books api error: %v", err)
		context.JSON(
			http.StatusBadGateway,
			gin.H{"error": "Call to google books api failed"},
		)
	}
	defer resp.Body.Close()

	// Decode the api response
	var apiResp *models.GoogleBooksAPIResponse
	apiResp, err = services.DecodeBookData(resp)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Error decoding api response"},
		)
	}

	// Convert to books
	var books []models.Book
	books, err = services.CreateBooks(apiResp)
	if err != nil {
		log.Printf("Error converting to book objects: %v", err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Error creating books list"},
		)
	}

	// Caching would go here
	context.JSON(http.StatusOK, books)
}
