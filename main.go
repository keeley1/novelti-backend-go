package main

import (
	"fmt"
	"log"

	// "net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/keeley1/novelti-backend-go/handlers"
)

/*
This project should follow best practices for a go server, use tests, good error handling
and have a good file structure.
*/

func main() {
	fmt.Println("hello world")

	// Should this move to server file
	// Create a default server
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	// Define the routes
	router.GET("/testsearch/:query", handlers.GetTestSearchHandler)
	router.GET("/booksbytitle/:title", handlers.GetBooksByTitleHandler)

	// Run the server
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
