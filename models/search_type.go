package models

type searchType string

const (
	SearchByTitle searchType = "title"
	SearchByGenre searchType = "genre"
	SearchBooks   searchType = "search"
)
