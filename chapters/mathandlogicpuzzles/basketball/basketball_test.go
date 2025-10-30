package basketball

import "testing"

func almostEqual(a, b float64) bool {
	const eps = 1e-9
	if a > b {
		return a-b < eps
	}
	return b-a < eps
}

func TestWinProbability(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		p        float64
		wantProb float64
	}{
		{"always miss", 3, 2, 0.0, 0.0},
		{"always hit", 3, 2, 1.0, 1.0},
		{"one shot p=0.5", 1, 1, 0.5, 0.5},
		{"two of three p=0.5", 3, 2, 0.5, 0.5}, // symmetric at p=0.5
		{"two of three p=0.7", 3, 2, 0.7, 0.784},
		{"two of three p=0.3", 3, 2, 0.3, 0.216},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WinProbability(tt.n, tt.k, tt.p)
			if !almostEqual(got, tt.wantProb) {
				t.Errorf("WinProbability(%d,%d,%.2f)=%.6f, want %.6f",
					tt.n, tt.k, tt.p, got, tt.wantProb)
			}
		})
	}
}

func TestGameComparison(t *testing.T) {
	// Find rough threshold where Game2 becomes better than Game1
	threshold := 0.0
	for p := 0.0; p <= 1.0001; p += 0.01 {
		g1 := p
		g2 := WinProbability(3, 2, p)
		if g2 > g1 {
			threshold = p
			break
		}
	}
	if threshold < 0.49 || threshold > 0.51 {
		t.Errorf("unexpected crossover threshold ≈ %.2f (expected ~0.38)", threshold)
	}
}
