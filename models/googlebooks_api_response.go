package models

// GoogleBooksAPIResponse is an API response object.
// It allows for safe unmarshalling of the API response.
// Extracts relevant book metadata.
type GoogleBooksAPIResponse struct {
	Items []struct {
		ID         string `json:"id"`
		VolumeInfo struct {
			Title         string   `json:"title"`
			Authors       []string `json:"authors"`
			PublishedDate string   `json:"publishedDate"`
			Description   string   `json:"description"`
			ImageLinks    struct {
				Thumbnail string `json:"thumbnail"`
			} `json:"imageLinks"`
		} `json:"volumeInfo"`
	} `json:"items"`
}
