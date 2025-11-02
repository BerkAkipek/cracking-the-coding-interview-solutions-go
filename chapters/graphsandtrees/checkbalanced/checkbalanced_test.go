package checkbalanced

import "testing"

// helper to build simple trees easily
func node(left, right *Node) *Node {
	return &Node{left: left, right: right}
}

func TestBalanceDFS(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want bool
	}{
		{
			name: "empty tree",
			root: nil,
			want: true, // empty is balanced
		},
		{
			name: "single node",
			root: &Node{},
			want: true,
		},
		{
			name: "perfectly balanced tree depth 2",
			root: node(&Node{}, &Node{}),
			want: true,
		},
		{
			name: "left-heavy but balanced",
			root: node(
				node(&Node{}, nil), // left child has one extra depth
				&Node{},
			),
			want: true,
		},
		{
			name: "right-heavy unbalanced",
			root: node(
				nil,
				node(nil, node(nil, &Node{})), // right subtree depth = 3
			),
			want: false,
		},
		{
			name: "deep left skewed chain",
			root: node(
				node(
					node(
						&Node{},
						nil,
					),
					nil,
				),
				nil,
			),
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BalanceDFS(tc.root)
			if got != tc.want {
				t.Errorf("BalanceDFS(%v) = %v; want %v", tc.name, got, tc.want)
			}
		})
	}
}
