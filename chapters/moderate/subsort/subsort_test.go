package subsort

import "testing"

func TestSubsort(t *testing.T) {
	tests := []struct {
		name      string
		arr       []int
		expectedL int
		expectedR int
	}{
		{
			name:      "CTCI corrected example (book errata fixed)",
			arr:       []int{1, 2, 4, 7, 10, 11, 1, 12, 6, 7, 16, 18, 19},
			expectedL: 1,
			expectedR: 9,
		},
		{
			name:      "Already sorted",
			arr:       []int{1, 2, 3, 4, 5},
			expectedL: -1,
			expectedR: -1,
		},
		{
			name:      "Single element",
			arr:       []int{42},
			expectedL: -1,
			expectedR: -1,
		},
		{
			name:      "Two elements sorted",
			arr:       []int{1, 2},
			expectedL: -1,
			expectedR: -1,
		},
		{
			name:      "Two elements reversed",
			arr:       []int{5, 1},
			expectedL: 0,
			expectedR: 1,
		},
		{
			name:      "Full reverse array",
			arr:       []int{9, 7, 5, 3, 1},
			expectedL: 0,
			expectedR: 4,
		},
		{
			name:      "Middle unsorted section",
			arr:       []int{1, 2, 8, 6, 7, 9, 10},
			expectedL: 2,
			expectedR: 4,
		},
		{
			name:      "Unsorted tail",
			arr:       []int{1, 2, 3, 10, 9, 8},
			expectedL: 3,
			expectedR: 5,
		},
		{
			name:      "Unsorted head",
			arr:       []int{10, 9, 8, 3, 4, 5},
			expectedL: 0,
			expectedR: 5,
		},
		{
			name:      "Duplicates inside window",
			arr:       []int{1, 3, 5, 4, 4, 6, 7},
			expectedL: 2,
			expectedR: 4,
		},
		{
			name:      "Multiple disorder regions",
			arr:       []int{1, 5, 3, 4, 2, 6, 7},
			expectedL: 1,
			expectedR: 4,
		},
		{
			name:      "All equal",
			arr:       []int{5, 5, 5, 5, 5},
			expectedL: -1,
			expectedR: -1,
		},
		{
			name:      "Random scattered disorder",
			arr:       []int{2, 6, 4, 8, 10, 9, 15},
			expectedL: 1,
			expectedR: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l, r := subsort(tc.arr)
			if l != tc.expectedL || r != tc.expectedR {
				t.Errorf("%s: subsort(%v) = (%d, %d); want (%d, %d)",
					tc.name, tc.arr, l, r, tc.expectedL, tc.expectedR)
			}
		})
	}
}
