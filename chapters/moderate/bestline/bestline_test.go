package bestline

import (
	"math"
	"testing"
)

func TestBestLine(t *testing.T) {
	tests := []struct {
		name     string
		points   []Point
		expected Line
		count    int
	}{
		{
			name: "Diagonal line y = x",
			points: []Point{
				{0, 0}, {1, 1}, {2, 2}, {3, 3},
			},
			expected: Line{M: 1, B: 0},
			count:    4,
		},
		{
			name: "Horizontal line y = 2",
			points: []Point{
				{0, 2}, {1, 2}, {2, 2}, {3, 5},
			},
			expected: Line{M: 0, B: 2},
			count:    3,
		},
		{
			name: "Vertical line x = 5",
			points: []Point{
				{5, 0}, {5, 2}, {5, 4}, {2, 2},
			},
			expected: Line{M: math.Inf(1), B: 5},
			count:    3,
		},
		{
			name: "Mixed lines, best is y = 1x + 0",
			points: []Point{
				{0, 0}, {1, 1}, {2, 2}, {2, 0}, {2, 4},
			},
			expected: Line{M: 1, B: 0},
			count:    3,
		},
		{
			name: "All points on different lines",
			points: []Point{
				{0, 0}, {1, 2}, {2, 5},
			},
			expected: Line{M: 2, B: 0},
			count:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLine, gotCount := BestLine(tt.points)

			if gotCount != tt.count {
				t.Errorf("%s: got count %d, want %d", tt.name, gotCount, tt.count)
			}

			if !almostEqual(gotLine.M, tt.expected.M) || !almostEqual(gotLine.B, tt.expected.B) {
				t.Errorf("%s: got line (M=%.6f, B=%.6f), want (M=%.6f, B=%.6f)",
					tt.name, gotLine.M, gotLine.B, tt.expected.M, tt.expected.B)
			}
		})
	}
}

func almostEqual(a, b float64) bool {
	if math.IsInf(a, 1) && math.IsInf(b, 1) {
		return true
	}
	return math.Abs(a-b) < 1e-6
}
