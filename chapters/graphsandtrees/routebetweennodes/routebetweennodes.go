package routebetweennodes

import (
	"crypto/sha256"
	"time"
)

/*
Route Between Nodes: Given a directed graph, design an algorithm to find out whether there is a
route between two nodes.
*/

type Node struct {
	ID    [32]byte
	Label string
	Data  any
}

func NewNode(label string, data any) *Node {
	id := sha256.Sum256([]byte(label + time.Now().String()))
	return &Node{
		ID:    id,
		Label: label,
		Data:  data,
	}
}

type Graph struct {
	Adj  map[[32]byte][][32]byte
	Root *Node
}

func NewGraph() *Graph {
	root := NewNode("root", []int{0})
	return &Graph{
		Adj:  map[[32]byte][][32]byte{},
		Root: root,
	}
}

func (g *Graph) Populate(from *Node, label string, data any) *Node {
	newNode := NewNode(label, data)
	g.Adj[from.ID] = append(g.Adj[from.ID], newNode.ID)
	return newNode
}

func (g *Graph) AddEdge(first, second *Node) bool {
	if first == nil || second == nil {
		return false
	}
	g.Adj[first.ID] = append(g.Adj[first.ID], second.ID)
	return true
}

type Queue struct {
	data [][32]byte
}

func (q *Queue) Enqueue(n [32]byte) bool {
	if n == [32]byte{} {
		return false
	}

	q.data = append(q.data, n)
	return true
}

func (q *Queue) Dequeue() ([32]byte, bool) {
	if len(q.data) == 0 {
		return [32]byte{}, false
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, true
}

func (g *Graph) HasRoute(node, target *Node) bool {
	if node == nil || target == nil {
		return false
	}

	visited := make(map[[32]byte]bool)
	queue := &Queue{
		data: [][32]byte{node.ID},
	}
	visited[node.ID] = true

	for len(queue.data) != 0 {
		currrent, ok := queue.Dequeue()
		if !ok {
			return false
		}

		for _, elem := range g.Adj[currrent] {
			if !visited[elem] {
				visited[elem] = true
				queue.Enqueue(elem)
			}
		}

		if currrent == target.ID {
			return true
		}
	}

	return false
}
