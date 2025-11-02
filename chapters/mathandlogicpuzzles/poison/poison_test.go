package poison

import (
	"testing"
)

func TestPoisonIdentification(t *testing.T) {
	const trials = 500
	for i := 0; i < trials; i++ {
		bottles, strips := Setup()

		EncodeStrips(bottles, strips)

		decoded := DecodeStrips(strips)

		var actual int
		for _, b := range bottles {
			if b.isPoisoned {
				actual = b.id
				break
			}
		}

		if decoded != actual {
			t.Fatalf("trial %d: decoded=%d, want=%d", i, decoded, actual)
		}
	}
}

func TestEncodeDoesNotMutateBottleIDs(t *testing.T) {
	bottles, strips := Setup()

	originalIDs := make([]int, len(bottles))
	for i, b := range bottles {
		originalIDs[i] = b.id
	}

	EncodeStrips(bottles, strips)

	for i, b := range bottles {
		if b.id != originalIDs[i] {
			t.Errorf("bottle %d id mutated: got %d, want %d", i, b.id, originalIDs[i])
		}
	}
}

func TestDecodeOutputRange(t *testing.T) {
	for i := 0; i < 100; i++ {
		bottles, strips := Setup()
		EncodeStrips(bottles, strips)
		id := DecodeStrips(strips)
		if id < 0 || id >= 1000 {
			t.Fatalf("decoded id %d out of range", id)
		}
	}
}

func TestStripLength(t *testing.T) {
	_, strips := Setup()
	if len(strips) != 10 {
		t.Fatalf("expected 10 strips, got %d", len(strips))
	}
}
