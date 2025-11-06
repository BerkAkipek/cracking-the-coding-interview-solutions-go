package numberswapper

import (
	"testing"
)

func TestSwapper(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		wantA int
		wantB int
	}{
		{"positive numbers", 5, 9, 9, 5},
		{"zero and positive", 0, 10, 10, 0},
		{"negative and positive", -7, 3, 3, -7},
		{"both negative", -4, -9, -9, -4},
		{"same value", 42, 42, 42, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := tt.a, tt.b
			Swapper(&a, &b)

			if a != tt.wantA || b != tt.wantB {
				t.Errorf("Swapper() failed for %s: got (a=%d, b=%d), want (a=%d, b=%d)",
					tt.name, a, b, tt.wantA, tt.wantB)
			}
		})
	}
}

// Optional: demonstrate what happens if both pointers are identical
func TestSwapper_SamePointer(t *testing.T) {
	x := 10
	Swapper(&x, &x)
	if x != 0 {
		t.Errorf("Swapper(&x, &x) should zero out x (got %d)", x)
	}
}
