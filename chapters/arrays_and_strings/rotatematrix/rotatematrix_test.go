package rotatematrix

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "nil matrix",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty matrix",
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			name: "1x1 matrix",
			input: [][]int{
				{42},
			},
			expected: [][]int{
				{42},
			},
		},
		{
			name: "2x2 matrix",
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "3x3 matrix",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "4x4 matrix",
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: [][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
		{
			name: "matrix with negatives",
			input: [][]int{
				{-1, -2, -3},
				{-4, -5, -6},
				{-7, -8, -9},
			},
			expected: [][]int{
				{-7, -4, -1},
				{-8, -5, -2},
				{-9, -6, -3},
			},
		},
		{
			name: "matrix with duplicates and zeros",
			input: [][]int{
				{0, 0, 1},
				{1, 0, 1},
				{1, 0, 0},
			},
			expected: [][]int{
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var inputCopy [][]int
			if tc.input != nil {
				inputCopy = make([][]int, len(tc.input))
				for i := range tc.input {
					inputCopy[i] = make([]int, len(tc.input[i]))
					copy(inputCopy[i], tc.input[i])
				}
			}

			RotateMatrix(inputCopy)

			if !reflect.DeepEqual(inputCopy, tc.expected) {
				t.Errorf("%s failed:\n got      = %v\n expected = %v", tc.name, inputCopy, tc.expected)
			}
		})
	}
}
