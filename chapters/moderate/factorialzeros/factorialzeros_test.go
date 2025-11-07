package factorialzeros

import "testing"

func TestTrailingZeros(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{3, 0},
		{5, 1},
		{10, 2},
		{25, 6},
		{50, 12},
		{100, 24},
		{125, 31},
	}

	for _, tt := range tests {
		got := TrailingZeros(tt.n)
		if got != tt.want {
			t.Errorf("TrailingZeros(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
