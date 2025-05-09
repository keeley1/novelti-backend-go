package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/models"
	"github.com/keeley1/novelti-backend-go/services"
	"github.com/keeley1/novelti-backend-go/utils"
)

func HandleBookSearch(context *gin.Context, searchQuery string, searchType string) {
	startIndexStr := context.DefaultQuery("startIndex", "0")
	startIndex := utils.ParseToPositiveInt(startIndexStr)

	resp, err := services.MakeAPICall(searchQuery, searchType, startIndex)
	if err != nil {
		log.Printf("Google books api error: %v", err)
		context.JSON(
			http.StatusBadGateway,
			gin.H{"error": "Call to google books api failed"},
		)
	}
	defer resp.Body.Close()

	// Decode the API response
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

	context.JSON(http.StatusOK, books)
}
