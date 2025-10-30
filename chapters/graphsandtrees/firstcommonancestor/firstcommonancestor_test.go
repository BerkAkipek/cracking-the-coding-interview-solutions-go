package firstcommonancestor

import "testing"

func TestFirstCommonAncestor(t *testing.T) {
	// Build the tree
	a := &Node{}
	b := &Node{}
	c := &Node{}
	d := &Node{}
	e := &Node{}
	f := &Node{}

	a.left, a.right = b, c
	b.left, b.right = d, e
	e.left = f

	tests := []struct {
		name     string
		x, y     *Node
		expected *Node
	}{
		{
			name:     "Common ancestor in left subtree (D,E)",
			x:        d,
			y:        e,
			expected: b,
		},
		{
			name:     "Common ancestor is root (D,C)",
			x:        d,
			y:        c,
			expected: a,
		},
		{
			name:     "One node is ancestor of other (B,F)",
			x:        b,
			y:        f,
			expected: b,
		},
		{
			name:     "Both nodes identical (E,E)",
			x:        e,
			y:        e,
			expected: e,
		},
		{
			name:     "One node not in tree",
			x:        d,
			y:        &Node{}, // new node not in tree
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstCommonAncestor(a, tt.x, tt.y)
			if got != tt.expected {
				t.Errorf("expected %p, got %p", tt.expected, got)
			}
		})
	}
}
