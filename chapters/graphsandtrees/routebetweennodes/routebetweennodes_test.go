package routebetweennodes

import "testing"

// helper function for quick linking
func link(from *Node, to ...*Node) *Node {
	from.Neighbours = append(from.Neighbours, to...)
	return from
}

func TestBFS(t *testing.T) {
	a := &Node{ID: 0}
	b := &Node{ID: 1}
	c := &Node{ID: 2}
	d := &Node{ID: 3}
	e := &Node{ID: 4}

	link(a, b, c)
	link(b, d)

	tests := []struct {
		name   string
		start  *Node
		target *Node
		want   bool
	}{
		// --- Positive cases ---
		{"A→D exists via B", a, d, true},
		{"A→C direct edge", a, c, true},
		{"B→D direct edge", b, d, true},
		{"A→A same node", a, a, true},

		// --- Negative cases ---
		{"B→A none (reverse edge)", b, a, false},
		{"C→D none", c, d, false},
		{"E→A isolated node", e, a, false},
		{"nil start", nil, a, false},
		{"nil target", a, nil, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := false
			if tc.start != nil {
				got = tc.start.BFS(tc.target)
			}
			if got != tc.want {
				t.Fatalf("BFS(%v→%v) = %v; want %v",
					nodeID(tc.start), nodeID(tc.target), got, tc.want)
			}
		})
	}
}

// helper for readable test names
func nodeID(n *Node) any {
	if n == nil {
		return "nil"
	}
	return n.ID
}
