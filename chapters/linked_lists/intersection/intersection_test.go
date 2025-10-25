package intersection

import "testing"

// --- Helper functions for building test lists ---

func makeChain(vals ...int) *Node {
	if len(vals) == 0 {
		return nil
	}
	head := &Node{Value: vals[0]}
	current := head
	for _, v := range vals[1:] {
		current.Next = &Node{Value: v}
		current = current.Next
	}
	return head
}

func buildListFromHead(head *Node) *Linkedlist {
	return &Linkedlist{Head: head}
}

func connectTailToNode(list *Linkedlist, nd *Node) {
	if list.Head == nil {
		list.Head = nd
		return
	}
	tail := list.GetTail()
	tail.Next = nd
}

// --- Table-driven tests for IsIntersection ---

func TestIsIntersection(t *testing.T) {
	tests := []struct {
		name     string
		build    func() (*Linkedlist, *Linkedlist, *Node)
		expected *Node
	}{
		{
			name: "Both lists empty",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				return &Linkedlist{}, &Linkedlist{}, nil
			},
			expected: nil,
		},
		{
			name: "One list empty",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				a := buildListFromHead(makeChain(1, 2, 3))
				b := &Linkedlist{}
				return a, b, nil
			},
			expected: nil,
		},
		{
			name: "Disjoint lists (different tails)",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				a := buildListFromHead(makeChain(1, 2, 3))
				b := buildListFromHead(makeChain(4, 5, 6))
				return a, b, nil
			},
			expected: nil,
		},
		{
			name: "Intersection at head (fully shared)",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				shared := makeChain(7, 8, 9)
				a := buildListFromHead(shared)
				b := buildListFromHead(shared)
				return a, b, shared
			},
			expected: func() *Node { shared := makeChain(7, 8, 9); return shared }(),
		},
		{
			name: "Intersection in the middle",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				shared := makeChain(7, 8, 9)
				a := buildListFromHead(makeChain(1, 2))
				b := buildListFromHead(makeChain(4, 5))
				connectTailToNode(a, shared)
				connectTailToNode(b, shared)
				return a, b, shared
			},
			expected: func() *Node { return makeChain(7, 8, 9) }(), // shared reference expected
		},
		{
			name: "Intersection at last node only",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				shared := makeChain(9)
				a := buildListFromHead(makeChain(1, 2))
				b := buildListFromHead(makeChain(3, 4))
				connectTailToNode(a, shared)
				connectTailToNode(b, shared)
				return a, b, shared
			},
			expected: func() *Node { return makeChain(9) }(),
		},
		{
			name: "Different lengths, intersecting",
			build: func() (*Linkedlist, *Linkedlist, *Node) {
				shared := makeChain(50, 60)
				a := buildListFromHead(makeChain(1, 2, 3))
				b := buildListFromHead(makeChain(9))
				connectTailToNode(a, shared)
				connectTailToNode(b, shared)
				return a, b, shared
			},
			expected: func() *Node { return makeChain(50, 60) }(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1, l2, expected := tt.build()
			got := IsIntersection(l1, l2)
			if got != expected {
				t.Errorf("%s: IsIntersection() = %v, expected %v", tt.name, got, expected)
			}
		})
	}
}
