package utils

import "testing"

func TestParseToPositiveInt(t *testing.T) {
	parsedResult := ParseToPositiveInt("10")

	if parsedResult != 10 {
		t.Errorf("Result was not correct, got: %v, want: %v.", parsedResult, 10)
	}
}

// Table driven test:
func TestParseToPositiveIntTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// Testing table
		{`"10" should be 10`, "10", 10},
		{`"-5" should be 0`, "-5", 0},
		{`"50" should be 50`, "50", 50},
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ParseToPositiveInt(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
