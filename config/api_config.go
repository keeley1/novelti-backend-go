// Package config contains application configuration constants
// such as API URLs and default query parameters.
package config

// GoogleBooksBaseURL is the base URL for the google books API call
const (
	GoogleBooksBaseURL = "https://www.googleapis.com/books/v1/volumes"
	DefaultOrderBy     = "newest"
	DefaultMaxResults  = 40
)
