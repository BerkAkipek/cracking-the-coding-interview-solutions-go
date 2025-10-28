package listofdepths

/*
List of Depths: Given a binary tree, design an algorithm which creates a linked list of all the nodes
at each depth (e.g., if you have a tree with depth D, you'll have D linked lists).
*/

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

type LinkedListNode struct {
	elements []*TreeNode
	next     *LinkedListNode
}

type Queue struct {
	data []*TreeNode
}

func (q *Queue) Enqueue(n *TreeNode) {
	q.data = append(q.data, n)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.data) == 0 {
		return nil
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element
}

func (n *TreeNode) ListOfDepths() *LinkedListNode {
	if n == nil {
		return nil
	}
	q := &Queue{
		data: []*TreeNode{n},
	}
	var head, tail *LinkedListNode

	for len(q.data) != 0 {
		levelSize := len(q.data)
		level := []*TreeNode{}
		for range levelSize {
			node := q.Dequeue()
			level = append(level, node)
			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
		}
		newNode := &LinkedListNode{elements: level}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
	return head
}
