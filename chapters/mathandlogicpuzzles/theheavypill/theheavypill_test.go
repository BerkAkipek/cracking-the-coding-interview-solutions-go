package theheavypill

import (
	"math"
	"testing"
)

func TestSumUntilN(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 3},
		{3, 6},
		{5, 15},
		{10, 55},
	}
	for _, tt := range tests {
		if got := sumUntilN(tt.n); got != tt.want {
			t.Errorf("sumUntilN(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestSetup(t *testing.T) {
	bottles := setup()

	if len(bottles) != 20 {
		t.Fatalf("expected 20 bottles, got %d", len(bottles))
	}

	countNormal := 0
	countSpecial := 0
	for _, b := range bottles {
		if len(b) != 20 {
			t.Errorf("bottle has %d pills, expected 20", len(b))
		}

		switch b[0] {
		case 1.0:
			countNormal++
		case 1.1:
			countSpecial++
		default:
			t.Errorf("unexpected pill weight %.1f", b[0])
		}
	}

	if countNormal != 19 {
		t.Errorf("expected 19 normal bottles, got %d", countNormal)
	}
	if countSpecial != 1 {
		t.Errorf("expected 1 special bottle, got %d", countSpecial)
	}
}

func TestShuffle(t *testing.T) {
	b1 := setup()
	b2 := setup()
	shuffle(b2)

	// Compute total weights regardless of order
	total1 := 0.0
	total2 := 0.0
	for _, b := range b1 {
		for _, p := range b {
			total1 += p
		}
	}
	for _, b := range b2 {
		for _, p := range b {
			total2 += p
		}
	}

	epsilon := 1e-9
	if math.Abs(total1-total2) > epsilon {
		t.Errorf("shuffle should not change total weight: before %.10f, after %.10f", total1, total2)
	}

	// Optional: verify that the order likely changed
	sameOrder := true
	for i := range b1 {
		for j := range b1[i] {
			if b1[i][j] != b2[i][j] {
				sameOrder = false
				break
			}
		}
		if !sameOrder {
			break
		}
	}
	if sameOrder {
		t.Errorf("shuffle likely did not change order")
	}
}
