package deckofcards

import (
	"math"
	"testing"
)

func approxEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

func TestMonteCarloConvergence(t *testing.T) {
	tests := []struct {
		name       string
		sampleSize int
		wantDealer float64
		wantPlayer float64
		tolerance  float64
	}{
		{"small 1e3", 1000, 0.50, 0.50, 0.10},
		{"medium 1e4", 10000, 0.50, 0.50, 0.05},
		{"large 1e5", 100000, 0.50, 0.50, 0.02},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dealer, player, tie := MonteCarlo(tc.sampleSize)
			sum := dealer + player + tie
			if !approxEqual(sum, 1.0, 0.02) {
				t.Fatalf("probabilities not normalized: got %.2f", sum)
			}

			diff := math.Abs(dealer - player)
			if diff > tc.tolerance {
				t.Fatalf("expected near-symmetric outcomes (±%.2f): dealer=%.2f player=%.2f",
					tc.tolerance, dealer, player)
			}

			if !approxEqual(dealer, tc.wantDealer, tc.tolerance) {
				t.Logf("dealer≈%.2f want≈%.2f (±%.2f)", dealer, tc.wantDealer, tc.tolerance)
			}
			if !approxEqual(player, tc.wantPlayer, tc.tolerance) {
				t.Logf("player≈%.2f want≈%.2f (±%.2f)", player, tc.wantPlayer, tc.tolerance)
			}
		})
	}
}
