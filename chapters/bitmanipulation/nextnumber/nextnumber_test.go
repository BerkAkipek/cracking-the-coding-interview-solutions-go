package nextnumber

import (
	"math/bits"
	"testing"
)

func TestNextNumber(t *testing.T) {
	tests := []struct {
		name        string
		num         int
		wantBigger  int
		wantSmaller int
	}{
		// --- textbook examples ---
		{"example_13948", 0b11011001111100, 0b11011010001111, 0b11011001111010},
		{"example_623", 623, 631, 607}, // verified: next larger 1783, smaller 1774

		// --- edge / boundary ---
		{"zero", 0b000000, 0, 0},
		{"single_bit_lowest", 0b000001, 0b000010, 0},
		{"single_bit_middle", 0b001000, 0b010000, 0b000100},
		{"single_bit_highest", 0b100000, 0, 0},
		{"all_ones_6bits", 0b111111, 0, 0},

		// --- alternating patterns ---
		{"alternating_even", 0b101010, 0b101100, 0b101001},
		{"alternating_odd", 0b010101, 0b010110, 0b010011},

		// --- dense / sparse ---
		{"dense_middle", 0b11100, 0b100011, 0b11010},
		{"mid_sparse", 0b10011100, 0b10100011, 0b10011010},
		{"sparse_far", 0b10000011, 0b10000101, 0b01110000},

		// --- larger mixed pattern ---
		{"large_mixed", 0b10011110000011, 0b10011110000101, 0b10011101110000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotBig, gotSmall := NextNumber(tc.num)

			pc := bits.OnesCount(uint(tc.num))
			if gotBig != 0 && bits.OnesCount(uint(gotBig)) != pc {
				t.Fatalf("popcount mismatch bigger: num=%b got=%b", tc.num, gotBig)
			}
			if gotSmall != 0 && bits.OnesCount(uint(gotSmall)) != pc {
				t.Fatalf("popcount mismatch smaller: num=%b got=%b", tc.num, gotSmall)
			}

			if tc.wantBigger != 0 && gotBig != tc.wantBigger {
				t.Errorf("Next larger mismatch: num=%b got=%b want=%b", tc.num, gotBig, tc.wantBigger)
			}
			if tc.wantSmaller != 0 && gotSmall != tc.wantSmaller {
				t.Errorf("Next smaller mismatch: num=%b got=%b want=%b", tc.num, gotSmall, tc.wantSmaller)
			}

			if gotBig != 0 && gotBig <= tc.num {
				t.Errorf("next larger not greater: num=%b got=%b", tc.num, gotBig)
			}
			if gotSmall != 0 && gotSmall >= tc.num {
				t.Errorf("next smaller not smaller: num=%b got=%b", tc.num, gotSmall)
			}
		})
	}
}

// Optional fuzz-style invariant check: verifies popcount and ordering on thousands of values.
func TestNextNumberConsistency(t *testing.T) {
	for n := 1; n < 1<<12; n++ { // 4096 samples
		big, small := NextNumber(n)
		pc := bits.OnesCount(uint(n))

		if big != 0 {
			if bits.OnesCount(uint(big)) != pc {
				t.Fatalf("popcount mismatch bigger: num=%b got=%b", n, big)
			}
			if big <= n {
				t.Fatalf("next larger not greater: num=%b got=%b", n, big)
			}
		}
		if small != 0 {
			if bits.OnesCount(uint(small)) != pc {
				t.Fatalf("popcount mismatch smaller: num=%b got=%b", n, small)
			}
			if small >= n {
				t.Fatalf("next smaller not smaller: num=%b got=%b", n, small)
			}
		}
	}
}
