// Package main is the entry point for the book application.
package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/handlers"
)

func main() {
	fmt.Println("hello world")

	// Create a default server
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	// Define the routes
	router.GET("/testsearch/:query", handlers.GetTestSearchHandler)
	router.GET("/booksbytitle/:title", handlers.GetBooksByTitleHandler)
	router.GET("/booksbygenre/:genre", handlers.GetBooksByGenreHandler)
	router.GET("/searchbooks/:searchquery", handlers.GetBooksBySearchHandler)

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
