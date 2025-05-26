// Package utils provides general utility functions and helpers
// used throughout the application to support common operations.
package utils

import (
	"fmt"
	"io"
)

// CloseBody safely closes resources.
func CloseBody(body io.Closer) error {
	if err := body.Close(); err != nil {
		return fmt.Errorf("failed to close body: %w", err)
	}
	return nil
}
