package jugofwater

import (
	"testing"
)

// helper to capture solution existence
func reachedGoal(A, B, T int) bool {
	if T > A && T > B || T%gcd(A, B) != 0 {
		return false
	}
	q := []S{{0, 0}}
	v := map[S]bool{}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if s.a == T || s.b == T {
			return true
		}
		if v[s] {
			continue
		}
		v[s] = true
		a, b := s.a, s.b
		next := []S{
			{A, b}, {a, B}, {0, b}, {a, 0},
			{a - min(a, B-b), b + min(a, B-b)},
			{a + min(b, A-a), b - min(b, A-a)},
		}
		q = append(q, next...)
	}
	return false
}

func TestBFSJugSolver(t *testing.T) {
	tests := []struct {
		name      string
		A, B, T   int
		wantReach bool
	}{
		// ✅ Feasible cases
		{"Classic_5_3_4", 5, 3, 4, true},
		{"Coprime_7_4_6", 7, 4, 6, true},
		{"EvenPair_10_6_8", 10, 6, 8, true},
		{"Coprime_9_4_7", 9, 4, 7, true},
		{"EqualJugs_4_4_4", 4, 4, 4, true},
		{"TargetZero_6_3_0", 6, 3, 0, true},

		// ❌ Infeasible (violates GCD)
		{"Infeasible_8_6_5", 8, 6, 5, false}, // gcd=2, target odd
		{"Infeasible_9_6_4", 9, 6, 4, false}, // gcd=3, 4 not multiple of 3
		{"Infeasible_12_8_7", 12, 8, 7, false},

		// ⚙️ Boundary
		{"GoalEqualsCapacityA", 5, 3, 5, true},
		{"GoalEqualsCapacityB", 5, 3, 3, true},
		{"GoalExceedsCapacity", 5, 3, 6, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := reachedGoal(tc.A, tc.B, tc.T)
			if got != tc.wantReach {
				t.Errorf("(%d,%d)->%d expected %v, got %v", tc.A, tc.B, tc.T, tc.wantReach, got)
			}
		})
	}
}