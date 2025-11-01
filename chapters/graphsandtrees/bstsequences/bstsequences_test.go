package bstsequences

import (
	"reflect"
	"sort"
	"testing"
)

// helper to normalize order-insensitive comparisons
func sort2D(slices [][]int) {
	sort.Slice(slices, func(i, j int) bool {
		a, b := slices[i], slices[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}

func TestBuildOrder(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected [][]int
	}{
		{
			name:     "SingleNode",
			values:   []int{2},
			expected: [][]int{{2}},
		},
		{
			name:     "SimpleBalanced",
			values:   []int{2, 1, 3},
			expected: [][]int{{2, 1, 3}, {2, 3, 1}},
		},
		{
			name:     "LeftSkewed",
			values:   []int{3, 2, 1},
			expected: [][]int{{3, 2, 1}},
		},
		{
			name:     "RightSkewed",
			values:   []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:   "BalancedFourNodes",
			values: []int{2, 1, 3, 4},
			// possible sequences that yield same BST:
			// root 2 → left 1 → right subtree [3,4] can appear in orders preserving 3 before 4
			expected: [][]int{
				{2, 1, 3, 4},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var root *Node
			for _, v := range tc.values {
				root = root.Insert(v)
			}
			got := root.BuildOrder()

			sort2D(got)
			sort2D(tc.expected)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("For %s:\n got  %v\n want %v", tc.name, got, tc.expected)
			}
		})
	}
}
