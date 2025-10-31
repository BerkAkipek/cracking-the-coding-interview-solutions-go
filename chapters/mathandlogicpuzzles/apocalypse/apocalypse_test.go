package apocalypse

import (
	"math"
	"testing"
)

func TestPopulationExperiment(t *testing.T) {
	tests := []struct {
		name        string
		numFamilies int
		tolerance   float64
	}{
		{"Small sample (100 families)", 100, 0.15},
		{"Medium sample (10_000 families)", 10_000, 0.05},
		{"Large sample (1_000_000 families)", 1_000_000, 0.01},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			totalRatio := 0.0
			numRuns := 10

			for range numRuns {
				totalRatio += PopulationExperiment(tc.numFamilies)
			}
			avgRatio := totalRatio / float64(numRuns)
			if math.Abs(avgRatio-0.5) > tc.tolerance {
				t.Errorf("expected ≈0.5, got %.4f (tolerance %.3f)", avgRatio, tc.tolerance)
			}
		})
	}
}
