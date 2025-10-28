package routebetweennodes

import "testing"

func TestQueue(t *testing.T) {
	q := &Queue{}
	id1 := [32]byte{1}
	id2 := [32]byte{2}

	if !q.Enqueue(id1) {
		t.Fatalf("Enqueue(%v) returned false unexpectedly", id1)
	}
	if !q.Enqueue(id2) {
		t.Fatalf("Enqueue(%v) returned false unexpectedly", id2)
	}

	got, ok := q.Dequeue()
	if !ok || got != id1 {
		t.Errorf("Dequeue() = %v, %v; want %v, true", got, ok, id1)
	}

	got, ok = q.Dequeue()
	if !ok || got != id2 {
		t.Errorf("Dequeue() = %v, %v; want %v, true", got, ok, id2)
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Dequeue() from empty queue should return ok=false")
	}
}

func TestAddEdgeAndPopulate(t *testing.T) {
	g := NewGraph()
	a := NewNode("A", nil)
	b := NewNode("B", nil)

	if ok := g.AddEdge(a, b); !ok {
		t.Errorf("AddEdge() returned false unexpectedly")
	}
	if len(g.Adj[a.ID]) != 1 || g.Adj[a.ID][0] != b.ID {
		t.Errorf("Adjacency list incorrect after AddEdge, got %v", g.Adj[a.ID])
	}

	// Populate should link from -> new node
	c := g.Populate(b, "C", nil)
	if len(g.Adj[b.ID]) == 0 || g.Adj[b.ID][0] != c.ID {
		t.Errorf("Populate() failed to link b -> c; got %v", g.Adj[b.ID])
	}
}

func TestHasRoute(t *testing.T) {
	type testCase struct {
		name     string
		build    func() (*Graph, *Node, *Node)
		expected bool
	}

	tests := []testCase{
		{
			name: "Direct connection (A -> B)",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				b := NewNode("B", nil)
				g.AddEdge(a, b)
				return g, a, b
			},
			expected: true,
		},
		{
			name: "Indirect connection (A -> B -> C)",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				b := g.Populate(a, "B", nil)
				c := g.Populate(b, "C", nil)
				return g, a, c
			},
			expected: true,
		},
		{
			name: "No connection (A -> B, C isolated)",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				b := NewNode("B", nil)
				c := NewNode("C", nil)
				g.AddEdge(a, b)
				return g, a, c
			},
			expected: false,
		},
		{
			name: "Cyclic graph (A -> B -> A)",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				b := NewNode("B", nil)
				g.AddEdge(a, b)
				g.AddEdge(b, a)
				return g, a, b
			},
			expected: true,
		},
		{
			name: "Self-loop (A -> A)",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				g.AddEdge(a, a)
				return g, a, a
			},
			expected: true,
		},
		{
			name: "Empty graph",
			build: func() (*Graph, *Node, *Node) {
				g := NewGraph()
				a := NewNode("A", nil)
				b := NewNode("B", nil)
				return g, a, b
			},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g, from, to := tc.build()
			got := g.HasRoute(from, to)
			if got != tc.expected {
				t.Errorf("HasRoute(%s) = %v; want %v", tc.name, got, tc.expected)
			}
		})
	}
}
