package validatebst

import "testing"

func n(v int, l, r *Node) *Node {
	return &Node{val: v, left: l, right: r}
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *Node
		want bool
	}{
		{"empty_tree", nil, true},
		{"single_node", n(42, nil, nil), true},
		{"perfect_valid_bst", n(8,
			n(4, n(2, nil, nil), n(6, nil, nil)),
			n(12, n(10, nil, nil), n(14, nil, nil))),
			true},
		{"deep_violation_right_subtree_smaller_than_root",
			n(10, n(5, nil, nil), n(15, n(6, nil, nil), n(20, nil, nil))),
			false},
		{"local_children_ok_but_global_violation",
			n(5, n(1, nil, nil), n(7, n(2, nil, nil), nil)),
			false},
		{"skewed_increasing_valid", n(1, nil, n(2, nil, n(3, nil, nil))), true},
		{"skewed_decreasing_valid", n(3, n(2, n(1, nil, nil), nil), nil), true},
		{"duplicates_on_left_invalid", n(5, n(5, nil, nil), nil), false},
		{"duplicates_on_right_invalid", n(5, nil, n(5, nil, nil)), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.root)
			if got != tc.want {
				t.Fatalf("IsValid() = %v; want %v", got, tc.want)
			}
		})
	}
}
