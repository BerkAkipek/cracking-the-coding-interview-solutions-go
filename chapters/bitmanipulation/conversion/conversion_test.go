package conversion

import "testing"

func TestConversion(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		// --- From problem statement ---
		{"example_29_15", 29, 15, 2}, // 11101 → 01111

		// --- Edge and sanity cases ---
		{"same_numbers", 42, 42, 0},     // no differing bits
		{"one_bit_difference", 1, 0, 1}, // 0001 → 0000
		{"flip_all_bits", 0, 255, 8},    // 00000000 → 11111111
		{"alternating_bits", 0b1010, 0b0101, 4},
		{"high_bits_difference", 1 << 31, 0, 1}, // only top bit differs (on 32-bit)
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Conversion(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Conversion(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
