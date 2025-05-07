package utils

import "testing"

// could also look into table driven tests
func TestParseToPositiveInt(t *testing.T) {
	parsedResult := ParseToPositiveInt("10")

	if parsedResult != 10 {
		t.Errorf("Result was not correct, got: %v, want: %v.", parsedResult, 10)
	}
}
