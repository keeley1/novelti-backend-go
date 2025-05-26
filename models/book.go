// Package models defines the data models to be used
// throughout the application.
package models

import (
	"html/template"
)

// Book is a simplified book model that can be returned to the client.
// It includes title, authors, published date, cover image, and ID metadata.
// Description field can safely render HTML content.
type Book struct {
	Title         string        `json:"title"`
	Authors       []string      `json:"authors"`
	PublishedDate string        `json:"publishedDate"`
	Cover         string        `json:"thumbnail"`
	ID            string        `json:"id"`
	Description   template.HTML `json:"description"`
}
