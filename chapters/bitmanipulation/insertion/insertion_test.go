package insertion

import (
	"fmt"
	"testing"
)

func TestInsertion(t *testing.T) {
	tests := []struct {
		name string
		n, m int32
		i, j int32
		want int32
	}{
		{
			name: "Example from book",
			n:    0b10000000000, // 1024
			m:    0b10011,       // 19
			i:    2,
			j:    6,
			want: 0b10001001100, // 1100 decimal
		},
		{
			name: "Insert into lower bits",
			n:    0b11111111111, // 2047
			m:    0b0,           // clear lower bits
			i:    0,
			j:    3,
			want: 0b11111110000, // bits 0..3 cleared
		},
		{
			name: "Insert single bit in middle",
			n:    0b10000000000, // 1024
			m:    0b1,
			i:    5,
			j:    5,
			want: 0b10000100000, // bit 5 set
		},
		{
			name: "Full overwrite (0..4)",
			n:    0b101010101010,
			m:    0b1111,
			i:    0,
			j:    3,
			want: (0b101010101010 & ^int32(0b1111)) | 0b1111,
		},
		{
			name: "Invalid range (i>j)",
			n:    0b1000,
			m:    0b1,
			i:    5,
			j:    2,
			want: 0, // function returns 0 on invalid input
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Insertion(tc.n, tc.m, tc.i, tc.j)
			if got != tc.want {
				t.Errorf("Insertion(%b,%b,%d,%d) = %b; want %b",
					tc.n, tc.m, tc.i, tc.j, got, tc.want)
			} else {
				fmt.Printf("%s: PASS => result: %b\n", tc.name, got)
			}
		})
	}
}
