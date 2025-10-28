package checkbalanced

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Empty tree",
			root:     nil,
			expected: true,
		},
		{
			name:     "Single node",
			root:     &TreeNode{},
			expected: true,
		},
		{
			name: "Perfectly balanced tree (3 levels)",
			root: &TreeNode{
				left: &TreeNode{
					left:  &TreeNode{},
					right: &TreeNode{},
				},
				right: &TreeNode{
					left:  &TreeNode{},
					right: &TreeNode{},
				},
			},
			expected: true,
		},
		{
			name: "Slightly unbalanced but valid",
			root: &TreeNode{
				left: &TreeNode{
					left: &TreeNode{},
				},
				right: &TreeNode{},
			},
			expected: true,
		},
		{
			name: "Clearly unbalanced (left heavy)",
			root: &TreeNode{
				left: &TreeNode{
					left: &TreeNode{
						left: &TreeNode{},
					},
				},
				right: &TreeNode{},
			},
			expected: false,
		},
		{
			name: "Clearly unbalanced (right heavy)",
			root: &TreeNode{
				left: &TreeNode{},
				right: &TreeNode{
					right: &TreeNode{
						right: &TreeNode{},
					},
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBalanced(tt.root)
			if result != tt.expected {
				t.Errorf("%s: got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
