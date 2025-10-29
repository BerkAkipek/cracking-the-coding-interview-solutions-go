package validatebst

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		root     *Node
		expected bool
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     &Node{val: 42},
			expected: true,
		},
		{
			name: "Simple valid BST",
			root: &Node{
				val:   8,
				left:  &Node{val: 4},
				right: &Node{val: 10},
			},
			expected: true,
		},
		{
			name: "Invalid left child greater than parent",
			root: &Node{
				val:   10,
				left:  &Node{val: 12},
				right: &Node{val: 15},
			},
			expected: false,
		},
		{
			name: "Invalid right child smaller than parent",
			root: &Node{
				val:   10,
				left:  &Node{val: 5},
				right: &Node{val: 7},
			},
			expected: false,
		},
		{
			name: "Deep invalid node (descendant violates range)",
			root: &Node{
				val: 10,
				left: &Node{
					val: 5,
					right: &Node{
						val: 11, // violates: > root but in left subtree
					},
				},
				right: &Node{val: 15},
			},
			expected: false,
		},
		{
			name: "Duplicates violate strict BST",
			root: &Node{
				val:   10,
				left:  &Node{val: 10}, // duplicate
				right: &Node{val: 12},
			},
			expected: false,
		},
		{
			name: "Valid tree with negative numbers",
			root: &Node{
				val:   0,
				left:  &Node{val: -5},
				right: &Node{val: 5},
			},
			expected: true,
		},
		{
			name: "Complex valid BST",
			root: &Node{
				val: 8,
				left: &Node{
					val:   4,
					left:  &Node{val: 2},
					right: &Node{val: 6},
				},
				right: &Node{
					val:   10,
					right: &Node{val: 20},
				},
			},
			expected: true,
		},
	}

	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.root, MinInt, MaxInt)
			if got != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}
