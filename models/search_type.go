package models

type searchType string

// SearchByTitle defines the search type where books are
// queried by title.
const (
	SearchByTitle searchType = "title"
	SearchByGenre searchType = "genre"
	SearchBooks   searchType = "search"
)
