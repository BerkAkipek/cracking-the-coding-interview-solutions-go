package antsonatriangle

import (
	"math"
	"testing"
)

func TestCollisionProbability(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected float64
	}{
		{"Triangle (n=3)", 3, 0.75},   // 1 - 1/2^(3-1)
		{"Square (n=4)", 4, 0.875},    // 1 - 1/2^(4-1)
		{"Pentagon (n=5)", 5, 0.9375}, // 1 - 1/2^(5-1)
		{"Hexagon (n=6)", 6, 0.96875}, // 1 - 1/2^(6-1)
		{"Single ant (n=1)", 1, 0.0},  // no collision possible
		{"Two ants (n=2)", 2, 0.5},    // 1 - 1/2^(2-1)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CollusionProbability(tt.n)
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("n=%d: got %.6f, expected %.6f", tt.n, got, tt.expected)
			}
		})
	}
}
