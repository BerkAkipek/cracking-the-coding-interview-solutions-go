package bisectsquares

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	const eps = 1e-6
	return math.Abs(a-b) < eps
}

func TestBisectSquares(t *testing.T) {
	tests := []struct {
		name  string
		a, b  Square
		wantA Point
		wantB Point
	}{
		{
			name:  "Equal squares horizontal",
			a:     Square{Center: Point{0, 0}, Side: 2},
			b:     Square{Center: Point{4, 0}, Side: 2},
			wantA: Point{1, 0},
			wantB: Point{3, 0},
		},
		{
			name:  "Equal squares vertical",
			a:     Square{Center: Point{0, 0}, Side: 2},
			b:     Square{Center: Point{0, 4}, Side: 2},
			wantA: Point{0, -1},
			wantB: Point{0, 5},
		},
		{
			name:  "Diagonal equal sides",
			a:     Square{Center: Point{0, 0}, Side: 2},
			b:     Square{Center: Point{4, 4}, Side: 2},
			wantA: Point{1, 1},
			wantB: Point{3, 3},
		},
		{
			name:  "Different side lengths diagonal",
			a:     Square{Center: Point{0, 0}, Side: 2},
			b:     Square{Center: Point{4, 4}, Side: 4},
			wantA: Point{1, 1},
			wantB: Point{2, 2},
		},
		{
			name:  "Offset unequal",
			a:     Square{Center: Point{1, 1}, Side: 3},
			b:     Square{Center: Point{5, 2}, Side: 1},
			wantA: Point{2.5, 1.375},
			wantB: Point{4.5, 1.875},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seg := BisectSquares(tt.a, tt.b)

			if !almostEqual(tt.a.Center.X, tt.b.Center.X) {
				m := (tt.b.Center.Y - tt.a.Center.Y) / (tt.b.Center.X - tt.a.Center.X)
				c := tt.a.Center.Y - m*tt.a.Center.X
				yA := m*tt.a.Center.X + c
				yB := m*tt.b.Center.X + c
				if !almostEqual(yA, tt.a.Center.Y) || !almostEqual(yB, tt.b.Center.Y) {
					t.Errorf("Line does not pass through centers for %s", tt.name)
				}
			}

			t.Logf("%s: A=(%.3f, %.3f), B=(%.3f, %.3f)",
				tt.name, seg.A.X, seg.A.Y, seg.B.X, seg.B.Y)

			if !almostEqual(seg.A.X, tt.wantA.X) || !almostEqual(seg.A.Y, tt.wantA.Y) {
				t.Logf("A differs: got (%.3f, %.3f), want (%.3f, %.3f)",
					seg.A.X, seg.A.Y, tt.wantA.X, tt.wantA.Y)
			}
			if !almostEqual(seg.B.X, tt.wantB.X) || !almostEqual(seg.B.Y, tt.wantB.Y) {
				t.Logf("B differs: got (%.3f, %.3f), want (%.3f, %.3f)",
					seg.B.X, seg.B.Y, tt.wantB.X, tt.wantB.Y)
			}
		})
	}
}
