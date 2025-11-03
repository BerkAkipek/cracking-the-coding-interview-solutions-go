package sortedsearch

import "testing"

func TestSortedSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "found in middle",
			arr:      []int{1, 3, 7, 10, 14, 25, 40, 90, 120, 160, 200, 250},
			target:   90,
			expected: 7,
		},
		{
			name:     "found at beginning",
			arr:      []int{1, 3, 7, 10, 14, 25},
			target:   1,
			expected: 0,
		},
		{
			name:     "found at end",
			arr:      []int{1, 3, 7, 10, 14, 25},
			target:   25,
			expected: 5,
		},
		{
			name:     "not found smaller than all",
			arr:      []int{5, 10, 15, 20},
			target:   1,
			expected: -1,
		},
		{
			name:     "not found larger than all",
			arr:      []int{5, 10, 15, 20},
			target:   30,
			expected: -1,
		},
		{
			name:     "array with duplicates",
			arr:      []int{1, 1, 1, 2, 2, 3, 3, 3},
			target:   2,
			expected: 3, // any valid index is fine
		},
		{
			name:     "single element present",
			arr:      []int{7},
			target:   7,
			expected: 0,
		},
		{
			name:     "single element absent",
			arr:      []int{7},
			target:   5,
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			target:   10,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedSearch(tt.arr, tt.target)
			if got != tt.expected {
				t.Errorf("SortedSearch(%v, %d) = %d, want %d",
					tt.arr, tt.target, got, tt.expected)
			}
		})
	}
}
