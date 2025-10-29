package successor

import (
	"testing"
)

func link(parent, left, right *Node) {
	if left != nil {
		parent.left = left
		left.parent = parent
	}
	if right != nil {
		parent.right = right
		right.parent = parent
	}
}

func TestSuccessor(t *testing.T) {
	n5 := &Node{}
	n17 := &Node{}
	n15 := &Node{right: n17}
	n10 := &Node{left: n5, right: n15}
	n25 := &Node{}
	n35 := &Node{}
	n30 := &Node{left: n25, right: n35}
	n20 := &Node{left: n10, right: n30}
	link(n20, n10, n30)
	link(n10, n5, n15)
	link(n15, nil, n17)
	link(n30, n25, n35)

	tests := []struct {
		name     string
		node     *Node
		expected *Node
	}{
		{"Successor of 5 is 10", n5, n10},
		{"Successor of 10 is 15", n10, n15},
		{"Successor of 15 is 17", n15, n17},
		{"Successor of 17 is 20", n17, n20},
		{"Successor of 20 is 25", n20, n25},
		{"Successor of 25 is 30", n25, n30},
		{"Successor of 30 is 35", n30, n35},
		{"Successor of 35 is nil", n35, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Successor(tt.node)
			if got != tt.expected {
				t.Errorf("%s: expected %p, got %p", tt.name, tt.expected, got)
			}
		})
	}
}
