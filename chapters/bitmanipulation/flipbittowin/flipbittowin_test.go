package flipbittowin

import "testing"

func TestFlipToWin(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		// --- From problem example ---
		{"example_1775", 1775, 8}, // 11011101111 → flip middle zero

		// --- Simple edge cases ---
		{"all_zeros", 0, 1},                  // flip one bit → 1
		{"single_one", 1, 2},                 // 0001 → flip neighbor zero → 11
		{"single_zero_between", 0b101, 3},    // 101 → flip middle zero → 111
		{"alternating_bits", 0b101010, 3},    // best is 3 via merging two 1s
		{"two_adjacent_zeros", 0b1100111, 4}, // only 4 possible, two zeros block full merge

		// --- Continuous ones ---
		{"all_ones_32bit", (1<<32 - 1), 32}, // already all ones → full length

		// --- Leading/trailing zeros ---
		{"leading_zero_run", 0b0001110111, 7},     // flip a zero to connect groups
		{"trailing_zero_run", 0b1110111000, 7},    // same idea at end
		{"multiple_zero_gaps", 0b111001110111, 7}, // largest connectable via one flip

		// --- Special numeric values ---
		{"max_int_32", 2147483647, 32}, // 0x7FFFFFFF → all but one bit set
		{"negative_number", -5, 0},     // negative guard (invalid)

		// --- Random patterns ---
		{"sparse_bits", 0b1000001, 2},       // flip one of zeros → 11
		{"dense_bits", 0b1110111101111, 9},  // connect 4+3+1 = 8
		{"long_gap", 0b111000000111, 4},     // only 1+1+1 = 4 possible
		{"scattered", 0b100111010011111, 6}, // best achievable merge is 6
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlipToWin(tt.num)
			if got != tt.want {
				t.Errorf("FlipToWin(%d / %b) = %d; want %d",
					tt.num, tt.num, got, tt.want)
			}
		})
	}
}
