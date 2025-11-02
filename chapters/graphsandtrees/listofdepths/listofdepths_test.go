package listofdepths

import (
	"reflect"
	"testing"
)

func n(val int, left, right *Node) *Node {
	return &Node{val: val, left: left, right: right}
}

func listToSlice(head *ListNode) []int {
	out := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		out = append(out, cur.Value)
	}
	return out
}

func printLevels(levels []*ListNode, t *testing.T) {
	for depth, head := range levels {
		values := listToSlice(head)
		t.Logf("Depth %d: %v", depth, values)
	}
}

func TestListOfDepths(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want [][]int
	}{
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "single node",
			root: n(1, nil, nil),
			want: [][]int{{1}},
		},
		{
			name: "balanced tree depth3",
			root: n(4,
				n(2, n(1, nil, nil), n(3, nil, nil)),
				n(6, n(5, nil, nil), n(7, nil, nil))),
			want: [][]int{
				{4},
				{2, 6},
				{1, 3, 5, 7},
			},
		},
		{
			name: "unbalanced left heavy",
			root: n(1,
				n(2,
					n(3,
						n(4, nil, nil),
						nil),
					nil),
				nil),
			want: [][]int{
				{1},
				{2},
				{3},
				{4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLists := ListOfDepths(tt.root)

			// special case: both nil
			if tt.root == nil {
				if gotLists != nil {
					t.Errorf("expected nil for empty tree, got %v", gotLists)
				}
				return
			}

			got := [][]int{}
			for _, head := range gotLists {
				got = append(got, listToSlice(head))
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListOfDepths() = %v, want %v", got, tt.want)
			}
		})
	}
}
