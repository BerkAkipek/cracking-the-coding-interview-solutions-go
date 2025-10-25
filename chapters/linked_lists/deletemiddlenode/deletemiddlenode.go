package deletemiddlenode

import "fmt"

/*
Delete Middle Node: Implement an algorithm to delete a node in the middle
(i.e., any node but the first and last node, not necessarily the exact middle)
of a singly linked list, given only access to that node.
EXAMPLE
lnput:the node c from the linked list: a->b->c->d->e->f
Result: nothing is returned, but the new linked list looks like a->b->d->e->f
*/

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{Value: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *LinkedList) Find(val int) *Node {
	current := l.Head
	for current != nil {
		if current.Value == val {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) ToSlice() []int {
	result := []int{}
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (l *LinkedList) NodeAt(ind int) *Node {
	if l.Head == nil || ind < 0 {
		return nil
	}

	current := l.Head
	for i := 0; current != nil && i < ind; i++ {
		current = current.Next
	}
	return current
}

func (l *LinkedList) PrintAll() {
	if l.Head == nil {
		fmt.Println("empty list")
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
}

func DeleteMiddleNode(n *Node) {
	if n.Next == nil {
		return
	}
	n.Value = n.Next.Value
	n.Next = n.Next.Next
}
