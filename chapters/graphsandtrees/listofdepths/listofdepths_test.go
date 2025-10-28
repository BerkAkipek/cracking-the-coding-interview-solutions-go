package listofdepths

import (
	"reflect"
	"testing"
)

// helper: connect nodes easily
func connect(parent, left, right *TreeNode) {
	parent.left = left
	parent.right = right
}

func TestListOfDepths(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		wantLens []int // number of elements per depth
	}{
		{
			name:     "Nil tree",
			root:     nil,
			wantLens: nil,
		},
		{
			name:     "Single node",
			root:     &TreeNode{},
			wantLens: []int{1},
		},
		{
			name: "Perfect binary tree of depth 3",
			root: func() *TreeNode {
				a := &TreeNode{}
				b := &TreeNode{}
				c := &TreeNode{}
				d := &TreeNode{}
				e := &TreeNode{}
				f := &TreeNode{}
				g := &TreeNode{}
				connect(a, b, c)
				connect(b, d, e)
				connect(c, f, g)
				return a
			}(),
			wantLens: []int{1, 2, 4},
		},
		{
			name: "Left skewed tree",
			root: func() *TreeNode {
				a := &TreeNode{}
				b := &TreeNode{}
				c := &TreeNode{}
				connect(a, b, nil)
				connect(b, c, nil)
				return a
			}(),
			wantLens: []int{1, 1, 1},
		},
		{
			name: "Right skewed tree",
			root: func() *TreeNode {
				a := &TreeNode{}
				b := &TreeNode{}
				c := &TreeNode{}
				connect(a, nil, b)
				connect(b, nil, c)
				return a
			}(),
			wantLens: []int{1, 1, 1},
		},
		{
			name: "Sparse tree (missing nodes at different levels)",
			root: func() *TreeNode {
				a := &TreeNode{}
				b := &TreeNode{}
				c := &TreeNode{}
				d := &TreeNode{}
				connect(a, b, nil)
				connect(b, nil, c)
				connect(c, d, nil)
				return a
			}(),
			wantLens: []int{1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.root.ListOfDepths()
			var gotLens []int
			for node := got; node != nil; node = node.next {
				gotLens = append(gotLens, len(node.elements))
			}

			if !reflect.DeepEqual(gotLens, tt.wantLens) {
				t.Errorf("ListOfDepths() got %v, want %v", gotLens, tt.wantLens)
			}
		})
	}
}
