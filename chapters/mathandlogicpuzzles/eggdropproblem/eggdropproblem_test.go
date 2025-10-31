package eggdropproblem

import (
	"math"
	"testing"
)

// TestTriangularSearch_NonDeterministic verifies that the randomized
// egg drop simulation always terminates with valid, in-range results.
// It also performs a probabilistic sanity check over many runs.
func TestTriangularSearch_NonDeterministic(t *testing.T) {
	tests := []struct {
		name        string
		numFloors   int
		repetitions int
	}{
		{"TinyBuilding_10Floors", 10, 200},
		{"SmallBuilding_50Floors", 50, 500},
		{"MediumBuilding_100Floors", 100, 500},
		{"HugeBuilding_500Floors", 500, 1000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			histogram := make([]int, tc.numFloors+1)

			for i := 0; i < tc.repetitions; i++ {
				result := TriangularSearch(tc.numFloors)
				// Invariant 1: must always return a valid floor
				if result < 1 || result > tc.numFloors {
					t.Fatalf("invalid result: %d (range 1–%d)", result, tc.numFloors)
				}
				histogram[result]++
			}

			// Invariant 2: results should cover a meaningful range of floors.
			distinct := 0
			for i := 1; i <= tc.numFloors; i++ {
				if histogram[i] > 0 {
					distinct++
				}
			}

			// Expected distinct ≈ n * (1 - (1 - 1/n)^k)
			expected := float64(tc.numFloors) *
				(1 - math.Pow(1-1/float64(tc.numFloors), float64(tc.repetitions)))

			// Allow 25% tolerance below expected.
			if float64(distinct) < 0.75*expected {
				t.Errorf("too few distinct results: got %d, expected ≈%.1f (floors=%d, trials=%d)",
					distinct, expected, tc.numFloors, tc.repetitions)
			}
		})
	}
}

// TestFindThreshold_Behavior verifies FindThreshold() directly
// with a deterministic predicate (no randomness) to confirm logic.
func TestFindThreshold_Behavior(t *testing.T) {
	cases := []struct {
		numFloors int
		breakAt   int
	}{
		{10, 1}, {10, 5}, {10, 10},
		{100, 37}, {100, 100},
		{500, 123},
	}

	for _, c := range cases {
		got := FindThreshold(c.numFloors, func(f int) bool { return f >= c.breakAt })
		if got != c.breakAt {
			t.Errorf("expected threshold %d, got %d (numFloors=%d)", c.breakAt, got, c.numFloors)
		}
	}
}
