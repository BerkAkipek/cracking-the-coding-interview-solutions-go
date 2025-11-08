package operations

import "testing"

func TestSubtraction(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive minus positive", 10, 3, 7},
		{"positive minus zero", 5, 0, 5},
		{"zero minus positive", 0, 7, -7},
		{"negative minus positive", -5, 3, -8},
		{"positive minus negative", 6, -2, 8},
		{"negative minus negative", -4, -6, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtraction(tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("Subtraction(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive * positive", 3, 4, 12},
		{"positive * zero", 8, 0, 0},
		{"zero * positive", 0, 5, 0},
		{"positive * negative", 3, -4, -12},
		{"negative * positive", -5, 2, -10},
		{"negative * negative", -3, -3, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiplication(tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("Multiplication(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name     string
		x, y     int
		expected int
	}{
		{"positive / positive", 10, 2, 5},
		{"positive / larger positive", 5, 10, 0},
		{"negative / positive", -9, 3, -3},
		{"positive / negative", 8, -2, -4},
		{"negative / negative", -12, -3, 4},
		{"zero / positive", 0, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Division(tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("Division(%d, %d) = %d; want %d", tt.x, tt.y, got, tt.expected)
			}
		})
	}
}

func TestDivisionByZeroPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Division did not panic on division by zero")
		}
	}()
	_ = Division(10, 0)
}
