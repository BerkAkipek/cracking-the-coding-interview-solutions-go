package numbermax

import (
	"math"
	"testing"
)

func TestCompareNums(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		// Basic positive numbers
		{"X greater", 10, 5, 10},
		{"Y greater", 3, 9, 9},
		{"Equal numbers", 8, 8, 8},

		// Negative numbers
		{"Both negative, X greater", -3, -7, -3},
		{"Both negative, Y greater", -10, -2, -2},

		// Mixed signs
		{"X negative, Y positive", -5, 2, 2},
		{"X positive, Y negative", 9, -1, 9},

		// Edge cases
		{"MaxInt vs MinInt", math.MaxInt64, math.MinInt64, math.MaxInt64},
		{"MinInt vs MaxInt", math.MinInt64, math.MaxInt64, math.MaxInt64},
		{"Zero vs Positive", 0, 5, 5},
		{"Zero vs Negative", 0, -9, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareNums(tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("CompareNums(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
			}
		})
	}
}
