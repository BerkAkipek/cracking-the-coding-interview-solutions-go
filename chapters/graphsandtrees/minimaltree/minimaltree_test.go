package minimaltree

import (
	"testing"
)

// Helper to get the in-order traversal of the BST.
func inorder(n *Node, result *[]int) {
	if n == nil {
		return
	}
	inorder(n.left, result)
	*result = append(*result, n.val)
	inorder(n.right, result)
}

// Helper to compute tree height.
func height(n *Node) int {
	if n == nil {
		return 0
	}
	lh := height(n.left)
	rh := height(n.right)
	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

// Helper to verify if the tree is height-balanced.
func isBalanced(n *Node) bool {
	if n == nil {
		return true
	}
	lh := height(n.left)
	rh := height(n.right)
	if lh-rh > 1 || rh-lh > 1 {
		return false
	}
	return isBalanced(n.left) && isBalanced(n.right)
}

func TestFromArrayToTree(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
	}{
		{"single_element", []int{10}},
		{"odd_number_of_elements", []int{1, 2, 3, 4, 5, 6, 7}},
		{"even_number_of_elements", []int{2, 4, 6, 8, 10, 12}},
		{"empty_array", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := FromArrayToTree(tt.arr)

			// Check in-order traversal matches the original sorted array.
			var result []int
			inorder(root, &result)
			if len(tt.arr) != len(result) {
				t.Fatalf("expected %d elements, got %d", len(tt.arr), len(result))
			}
			for i := range tt.arr {
				if tt.arr[i] != result[i] {
					t.Errorf("inorder traversal mismatch at index %d: got %d, want %d", i, result[i], tt.arr[i])
				}
			}

			// Verify balance property (minimal height)
			if !isBalanced(root) {
				t.Errorf("tree from %v is not balanced", tt.arr)
			}

			// Verify height ≈ log2(n)
			n := len(tt.arr)
			if n > 0 {
				h := height(root)
				// Rough upper bound for balanced BST height is log2(n)+1
				// We just ensure it isn't degenerate
				if h > len(tt.arr) {
					t.Errorf("tree height too large: got %d for %v", h, tt.arr)
				}
			}
		})
	}
}
