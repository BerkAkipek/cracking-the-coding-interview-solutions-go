package dominos

import "testing"

func TestCanTile(t *testing.T) {
	tests := []struct {
		name     string
		removed  []Cell
		wantTile bool
		wantB    int
		wantW    int
	}{
		{
			name:     "no removed squares",
			removed:  nil,
			wantTile: true, // 32 vs 32
			wantB:    32,
			wantW:    32,
		},
		{
			name:     "diagonally opposite corners removed",
			removed:  []Cell{{0, 0}, {7, 7}},
			wantTile: false, // 30 vs 32
			wantB:    30,
			wantW:    32,
		},
		{
			name:     "adjacent corners removed",
			removed:  []Cell{{0, 0}, {0, 1}},
			wantTile: true, // 31 vs 31
			wantB:    31,
			wantW:    31,
		},
		{
			name:     "random same-color removals",
			removed:  []Cell{{1, 1}, {3, 3}}, // both white
			wantTile: false,
			wantB:    30,
			wantW:    32,
		},
		{
			name:     "random opposite-color removals",
			removed:  []Cell{{1, 1}, {2, 1}}, // one white, one black
			wantTile: true,
			wantB:    31,
			wantW:    31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTile, gotB, gotW := CanTile(tt.removed)
			if gotTile != tt.wantTile {
				t.Errorf("CanTile() = %v, want %v", gotTile, tt.wantTile)
			}
			if gotB != tt.wantB || gotW != tt.wantW {
				t.Errorf("color counts = (B=%d,W=%d), want (B=%d,W=%d)",
					gotB, gotW, tt.wantB, tt.wantW)
			}
		})
	}
}
