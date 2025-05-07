package utils

import (
	"strconv"
	"strings"
)

func ParseToPositiveInt(queryParam string) int {
	cleansedQueryParam, err := strconv.Atoi(strings.TrimSpace(queryParam))
	if err != nil || cleansedQueryParam < 0 {
		cleansedQueryParam = 0
	}
	return cleansedQueryParam
}
