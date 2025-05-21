package utils

import (
	"strconv"
	"strings"
)

// ParseToPositiveInt parses a given string to a positive number.
// Returns 0 is the number is invalid (e.g. negative).
func ParseToPositiveInt(queryParam string) int {
	cleansedQueryParam, err := strconv.Atoi(strings.TrimSpace(queryParam))
	if err != nil || cleansedQueryParam < 0 {
		cleansedQueryParam = 0
	}
	return cleansedQueryParam
}
