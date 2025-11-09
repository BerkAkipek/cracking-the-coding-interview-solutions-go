package mastermind

import (
	"reflect"
	"testing"
)

func TestMasterMind_Counts(t *testing.T) {
	type exp struct{ hits, pseudos int }

	tests := []struct {
		name     string
		computer []string
		player   []string
		want     exp
	}{
		{
			name:     "all hits",
			computer: []string{"R", "G", "B", "Y"},
			player:   []string{"R", "G", "B", "Y"},
			want:     exp{hits: 4, pseudos: 0},
		},
		{
			name:     "all pseudo-hits (rotation)",
			computer: []string{"R", "G", "B", "Y"},
			player:   []string{"G", "B", "Y", "R"},
			want:     exp{hits: 0, pseudos: 4},
		},
		{
			name:     "no matches",
			computer: []string{"R", "R", "R", "R"},
			player:   []string{"G", "G", "G", "G"},
			want:     exp{hits: 0, pseudos: 0},
		},
		{
			name:     "duplicates, mixed hits and pseudo-hits (per current logic)",
			computer: []string{"R", "R", "G", "B"},
			player:   []string{"R", "G", "R", "B"},
			want:     exp{hits: 2, pseudos: 2},
		},
		{
			name:     "some hits some pseudos",
			computer: []string{"R", "G", "B", "Y"},
			player:   []string{"R", "B", "G", "Y"},
			want:     exp{hits: 2, pseudos: 2},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := MasterMind(tt.computer, tt.player)
			hits, pseudos := countResults(got)
			if hits != tt.want.hits || pseudos != tt.want.pseudos {
				t.Fatalf("got (hits=%d, pseudos=%d), want (hits=%d, pseudos=%d); raw=%v",
					hits, pseudos, tt.want.hits, tt.want.pseudos, got)
			}
		})
	}
}

func TestMasterMind_OrderMatchesPositions(t *testing.T) {
	computer := []string{"R", "G", "B", "Y"}
	player := []string{"R", "B", "G", "Y"}
	want := []string{"hit", "pseudo-hit", "pseudo-hit", "hit"}

	got := MasterMind(computer, player)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("order mismatch: got %v, want %v", got, want)
	}
}

func TestSetup_Sane(t *testing.T) {
	comp, player := Setup()
	if len(comp) != 4 || len(player) != 4 {
		t.Fatalf("expected both sequences of length 4, got comp=%d player=%d", len(comp), len(player))
	}
	for i := 0; i < 4; i++ {
		if !isAllowedColor(comp[i]) || !isAllowedColor(player[i]) {
			t.Fatalf("unexpected color at index %d: comp=%q player=%q", i, comp[i], player[i])
		}
	}
}

func countResults(res []string) (hits, pseudos int) {
	for _, r := range res {
		switch r {
		case "hit":
			hits++
		case "pseudo-hit":
			pseudos++
		}
	}
	return
}

func isAllowedColor(c string) bool {
	switch c {
	case "R", "G", "B", "Y":
		return true
	default:
		return false
	}
}
