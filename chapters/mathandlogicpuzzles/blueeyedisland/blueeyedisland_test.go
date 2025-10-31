package blueeyedisland

import (
	"testing"
)

// Helper: deterministic island creation (no randomness)
func createFixedIsland(total, numBlue int) *Island {
	people := []*Person{}
	for i := 0; i < total; i++ {
		color := "brown"
		if i < numBlue {
			color = "blue"
		}
		p := &Person{
			id:       i,
			eyeColor: color,
			isBlue:   (color == "blue"),
		}
		people = append(people, p)
	}
	return &Island{people: people}
}

// Modified Game logic for test returns days
func runSimulation(island *Island) int {
	numBlue := 0
	for _, p := range island.people {
		if p.eyeColor == "blue" {
			numBlue++
		}
	}
	for _, p := range island.people {
		if p.eyeColor == "blue" {
			p.visibleBlues = numBlue - 1
		} else {
			p.visibleBlues = numBlue
		}
	}

	day := 0
	for {
		day++
		leftToday := 0
		for _, p := range island.people {
			if p.departed {
				continue
			}
			if p.isBlue && day == p.visibleBlues+1 {
				p.departed = true
				leftToday++
			}
		}
		if leftToday == numBlue {
			return day
		}
	}
}

func TestBlueEyedDeparture(t *testing.T) {
	tests := []struct {
		name    string
		total   int
		numBlue int
		wantDay int
	}{
		{"OneBlue", 5, 1, 1},
		{"TwoBlue", 5, 2, 2},
		{"ThreeBlue", 5, 3, 3},
		{"TenBlue", 10, 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			island := createFixedIsland(tt.total, tt.numBlue)
			gotDay := runSimulation(island)

			if gotDay != tt.wantDay {
				t.Errorf("expected day %d, got %d", tt.wantDay, gotDay)
			}

			// Ensure all blue-eyed departed
			for _, p := range island.people {
				if p.isBlue && !p.departed {
					t.Errorf("blue-eyed person %d did not depart", p.id)
				}
			}
		})
	}
}
