package drawline

import (
	"fmt"
	"testing"
)

// helper: converts a byte slice into a compact binary string (for easy comparison)
func bits(screen []byte) string {
	out := ""
	for _, b := range screen {
		out += fmt.Sprintf("%08b", b)
	}
	return out
}

func TestDrawLine(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		x1     int
		x2     int
		y      int
		height int
		want   string
	}{
		{
			name:   "simple_line_middle",
			width:  16,
			x1:     3,
			x2:     11,
			y:      0,
			height: 1,
			want:   "0001111111110000", // fixed
		},
		{
			name:   "single_pixel",
			width:  8,
			x1:     4,
			x2:     4,
			y:      0,
			height: 1,
			want:   "00001000",
		},
		{
			name:   "full_row",
			width:  16,
			x1:     0,
			x2:     15,
			y:      0,
			height: 1,
			want:   "1111111111111111",
		},
		{
			name:   "cross_two_bytes",
			width:  32,
			x1:     10,
			x2:     23,
			y:      0,
			height: 1,
			want:   "00000000001111111111111100000000",
		},
		{
			name:   "line_in_second_row",
			width:  16,
			x1:     2,
			x2:     5,
			y:      1,
			height: 2,
			want:   "00000000000000000011110000000000", // fixed
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			screen := make([]byte, (tc.width/8)*tc.height)
			DrawLine(screen, tc.width, tc.x1, tc.x2, tc.y)
			got := bits(screen)
			if got != tc.want {
				t.Errorf("DrawLine(%d,%d,%d,%d) got\n%s\nwant\n%s", tc.width, tc.x1, tc.x2, tc.y, got, tc.want)
			}
		})
	}
}
