package listofdepths

/*
List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes
at each depth (e.g., if you have a tree with depth D, you'll have D linked lists)
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

type ListNode struct {
	Value int
	Next  *ListNode
}

type Queue struct {
	data []*Node
}

func (q *Queue) Enqueue(n *Node) {
	q.data = append(q.data, n)
}

func (q *Queue) Length() int {
	return len(q.data)
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

func ListOfDepths(root *Node) []*ListNode {
	if root == nil {
		return nil
	}

	q := NewQueue()
	result := []*ListNode{}

	q.Enqueue(root)
	for !q.IsEmpty() {
		level := q.Length()
		var head, tail *ListNode = nil, nil

		for range level {
			node := q.Dequeue()
			newNode := &ListNode{Value: node.val}

			if head == nil {
				head = newNode
				tail = newNode
			} else {
				tail.Next = newNode
				tail = newNode
			}

			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
		}
		result = append(result, head)
	}

	return result
}
