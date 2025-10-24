package zeromatrix

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "No zeros (matrix unchanged)",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "Single zero in middle",
			input: [][]int{
				{1, 2, 3},
				{4, 0, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 0, 3},
				{0, 0, 0},
				{7, 0, 9},
			},
		},
		{
			name: "Zero in first row and column",
			input: [][]int{
				{0, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 5, 6},
				{0, 8, 9},
			},
		},
		{
			name: "All zeros (everything becomes zero)",
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "Single element non-zero",
			input: [][]int{
				{5},
			},
			expected: [][]int{
				{5},
			},
		},
		{
			name: "Single element zero",
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			name: "Rectangular 2x3 matrix with zero",
			input: [][]int{
				{1, 2, 0},
				{4, 5, 6},
			},
			expected: [][]int{
				{0, 0, 0},
				{4, 5, 0},
			},
		},
		{
			name: "Rectangular 3x2 matrix with zero",
			input: [][]int{
				{1, 0},
				{3, 4},
				{5, 6},
			},
			expected: [][]int{
				{0, 0},
				{3, 0},
				{5, 0},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ZeroMatrix(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("ZeroMatrix(%v) = %v; expected: %v", tc.input, got, tc.expected)
			}
		})
	}
}
