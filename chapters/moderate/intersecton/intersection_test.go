package intersecton

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool {
	return math.Abs(a-b) < eps
}

func TestPointIntersection(t *testing.T) {
	tests := []struct {
		name   string
		a, b   [][]float64
		expect []float64 // nil → expect no intersection
	}{
		{
			name:   "simple X-shaped intersection",
			a:      [][]float64{{0, 0}, {4, 4}},
			b:      [][]float64{{0, 4}, {4, 0}},
			expect: []float64{2, 2},
		},
		{
			name:   "parallel horizontal lines",
			a:      [][]float64{{0, 0}, {4, 0}},
			b:      [][]float64{{0, 1}, {4, 1}},
			expect: nil,
		},
		{
			name:   "collinear non-overlapping",
			a:      [][]float64{{0, 0}, {1, 0}},
			b:      [][]float64{{2, 0}, {3, 0}},
			expect: nil,
		},
		{
			name:   "collinear points",
			a:      [][]float64{{0, 0}, {2, 0}},
			b:      [][]float64{{2, 0}, {4, 0}},
			expect: nil,
		},
		{
			name:   "intersection inside segment bounds",
			a:      [][]float64{{1, 1}, {4, 4}},
			b:      [][]float64{{1, 4}, {4, 1}},
			expect: []float64{2.5, 2.5},
		},
		{
			name:   "disjoint non-parallel segments",
			a:      [][]float64{{0, 0}, {1, 1}},
			b:      [][]float64{{2, 2}, {3, 3}},
			expect: nil,
		},
		{
			name:   "endpoint intersection",
			a:      [][]float64{{0, 0}, {2, 2}},
			b:      [][]float64{{2, 2}, {4, 0}},
			expect: []float64{2, 2},
		},
		{
			name:   "vertical and horizontal cross",
			a:      [][]float64{{1, -1}, {1, 1}},
			b:      [][]float64{{0, 0}, {2, 0}},
			expect: []float64{1, 0},
		},
		{
			name:   "diagonal non-intersecting far away",
			a:      [][]float64{{0, 0}, {1, 1}},
			b:      [][]float64{{2, 0}, {3, -1}},
			expect: nil,
		},
	}

	const eps = 1e-6
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PointIntersection(tc.a, tc.b)

			if tc.expect == nil {
				if got != nil {
					t.Errorf("expected no intersection, got %v", got)
				}
				return
			}

			if got == nil {
				t.Errorf("expected intersection at %v, got nil", tc.expect)
				return
			}

			if !almostEqual(got[0], tc.expect[0], eps) ||
				!almostEqual(got[1], tc.expect[1], eps) {
				t.Errorf("expected intersection at %v, got %v", tc.expect, got)
			}
		})
	}
}
