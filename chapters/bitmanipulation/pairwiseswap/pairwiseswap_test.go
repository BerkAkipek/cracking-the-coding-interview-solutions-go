package pairwiseswap

import "testing"

func TestPairwise64(t *testing.T) {
	tests := []struct {
		name string
		in   uint64
		want uint64
	}{
		// --- Alternating patterns ---
		{"alternating_AA", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555},
		{"alternating_55", 0x5555555555555555, 0xAAAAAAAAAAAAAAAA},

		// --- Small examples ---
		{"simple_alternating", 0b101010, 0b010101},
		{"reverse_alternating", 0b010101, 0b101010},
		{"zero", 0, 0},
		{"one", 1, 2},
		{"two", 2, 1},
		{"three", 3, 3},
		{"five", 5, 10},
		{"ten", 10, 5},

		// --- 32-bit visible patterns ---
		{"mix_pattern_32", 0x12345678, 0x2138A9B4},  
		{"mirror_pattern_32", 0x48D2AC3C, 0x84E15C3C},

		// --- 64-bit high patterns ---
		{"pattern_highbits_1", 0xF0F0F0F0F0F0F0F0, 0xF0F0F0F0F0F0F0F0},
		{"pattern_highbits_2", 0x0F0F0F0F0F0F0F0F, 0x0F0F0F0F0F0F0F0F},

		// --- Random 64-bit examples ---
		{"random_1", 0x123456789ABCDEF0, 0x2138A9B4657CEDF0}, 
		{"random_2", 0xFEDCBA9876543210, 0xFDEC7564B9A83120},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Pairwise(tc.in)
			if got != tc.want {
				t.Errorf("Pairwise(%#016x) = %#016x; want %#016x",
					tc.in, got, tc.want)
			}
		})
	}
}
