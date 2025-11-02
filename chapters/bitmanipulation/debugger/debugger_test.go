package debugger

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		// --- base cases ---
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},

		// --- larger powers of two ---
		{8, true},
		{16, true},
		{32, true},
		{64, true},

		// --- non-powers ---
		{6, false},
		{7, false},
		{9, false},
		{12, false},

		// --- negative numbers ---
		{-1, false},
		{-8, false},
	}

	for _, tc := range tests {
		got := IsPowerOfTwo(tc.n)
		if got != tc.want {
			t.Errorf("IsPowerOfTwo(%d) = %v; want %v", tc.n, got, tc.want)
		}
	}
}

func TestNextPowerOfTwo(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 4},
		{5, 8},
		{6, 8},
		{7, 8},
		{8, 8},
		{9, 16},
		{15, 16},
		{16, 16},
		{17, 32},
		{33, 64},
		{65, 128},
		{255, 256},
	}

	for _, tc := range tests {
		got := NextPowerOfTwo(tc.n)
		if got != tc.want {
			t.Errorf("NextPowerOfTwo(%d) = %d; want %d", tc.n, got, tc.want)
		}
	}
}

func TestPrevPowerOfTwo(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 4},
		{5, 4},
		{6, 4},
		{7, 4},
		{8, 8},
		{9, 8},
		{15, 8},
		{16, 16},
		{17, 16},
		{31, 16},
		{32, 32},
		{33, 32},
		{127, 64},
		{128, 128},
		{255, 128},
	}

	for _, tc := range tests {
		got := PrevPowerOfTwo(tc.n)
		if got != tc.want {
			t.Errorf("PrevPowerOfTwo(%d) = %d; want %d", tc.n, got, tc.want)
		}
	}
}
