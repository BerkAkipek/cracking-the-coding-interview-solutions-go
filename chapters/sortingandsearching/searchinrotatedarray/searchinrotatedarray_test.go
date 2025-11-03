package searchinrotatedarray

import "testing"

func TestSearchRotatedArray(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{[]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5, 8},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
	}

	for _, tt := range tests {
		if got := SearchRotatedArray(tt.arr, tt.target); got != tt.expected {
			t.Errorf("SearchRotatedArray(%v, %d) = %d; want %d", tt.arr, tt.target, got, tt.expected)
		}
	}
}
