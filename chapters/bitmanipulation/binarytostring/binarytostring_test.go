package binarytostring

import "testing"

func TestBinaryRepresentation(t *testing.T) {
	tests := []struct {
		name string
		num  float64
		want string
	}{
		{"half", 0.5, "0.1"},
		{"quarter", 0.25, "0.01"},
		{"three_quarters", 0.75, "0.11"},
		{"five_eighths", 0.625, "0.101"},
		{"almost_one", 0.999999, "ERROR"},
		{"ten_percent", 0.1, "ERROR"},
		{"seventy_two_percent", 0.72, "ERROR"},
		{"zero", 0.0, "ERROR"},
		{"one", 1.0, "ERROR"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BinaryRepresentation(tc.num)
			if got != tc.want {
				t.Errorf("BinaryRepresentation(%v) = %v; want %v", tc.num, got, tc.want)
			}
		})
	}
}
