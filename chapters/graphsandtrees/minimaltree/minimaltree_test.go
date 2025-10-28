package minimaltree

import (
	"math"
	"reflect"
	"testing"
)

func inorderTraversal(n *Node, result *[]int) {
	if n == nil {
		return
	}
	inorderTraversal(n.left, result)
	*result = append(*result, n.data)
	inorderTraversal(n.right, result)
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	left := height(n.left)
	right := height(n.right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{"Empty array", []int{}, []int{}},
		{"Single element", []int{10}, []int{10}},
		{"Odd number of elements", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"Even number of elements", []int{2, 4, 6, 8, 10, 12}, []int{2, 4, 6, 8, 10, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.arr) == 0 {
				root := BuildTree(tt.arr, 0, len(tt.arr)-1)
				if root != nil {
					t.Errorf("expected nil, got non-nil root")
				}
				return
			}

			root := BuildTree(tt.arr, 0, len(tt.arr)-1)
			var result []int
			inorderTraversal(root, &result)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("inorder traversal = %v, expected %v", result, tt.expected)
			}

			h := height(root)
			expectedMax := int(math.Floor(math.Log2(float64(len(tt.arr)))) + 1)
			if h > expectedMax {
				t.Errorf("tree too tall: got %d, expected ≤ %d", h, expectedMax)
			}
		})
	}
}

func TestAddElement(t *testing.T) {
	tests := []struct {
		name     string
		initial  []int
		toAdd    []int
		expected []int
	}{
		{"Insert into empty tree", []int{}, []int{5}, []int{5}},
		{"Insert one value", []int{3}, []int{1}, []int{1, 3}},
		{"Insert multiple values", []int{2, 4, 6}, []int{1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := &BinaryTree{}
			if len(tt.initial) > 0 {
				tree.Root = BuildTree(tt.initial, 0, len(tt.initial)-1)
			}

			for _, v := range tt.toAdd {
				tree.Root = insert(tree.Root, v)
			}

			var result []int
			inorderTraversal(tree.Root, &result)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("inorder traversal = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestBSTProperty(t *testing.T) {
	root := &BinaryTree{}
	values := []int{4, 2, 6, 1, 3, 5, 7}
	for _, v := range values {
		root.Root = insert(root.Root, v)
	}

	var result []int
	inorderTraversal(root.Root, &result)

	for i := 1; i < len(result); i++ {
		if result[i-1] > result[i] {
			t.Errorf("BST property violated at %v > %v", result[i-1], result[i])
		}
	}
}
