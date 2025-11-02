package routebetweennodes

/*
Route Between Nodes: Given a directed graph, design an algorithm to find out whether there is a
route between two nodes
*/

type Node struct {
	ID         int
	Neighbours []*Node
}

type Queue struct {
	data []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.data = append(q.data, n)
}

func (q *Queue) Dequeue() *Node {
	if len(q.data) == 0 {
		return nil
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}

func NewQueue() *Queue {
	return &Queue{
		data: []*Node{},
	}
}

func (q *Queue) IsEmpty() bool { return len(q.data) == 0 }

func (n *Node) BFS(other *Node) bool {
	q := NewQueue()
	seen := make(map[*Node]bool)

	q.Enqueue(n)
	seen[n] = true

	for !q.IsEmpty() {
		current := q.Dequeue()
		if current == other {
			return true
		}
		for _, nb := range current.Neighbours {
			if !seen[nb] {
				seen[nb] = true
				q.Enqueue(nb)
			}
		}
	}
	return false
}
