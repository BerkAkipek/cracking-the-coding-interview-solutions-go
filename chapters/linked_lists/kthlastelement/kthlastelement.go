package kthlastelement

import (
	"fmt"
)

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

func (l *LinkedList) Remove(val int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == val {
		l.Head = l.Head.Next
		return
	}

	prev, current := l.Head, l.Head.Next
	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}
}

func (l *LinkedList) PrintList() {
	if l.Head == nil {
		fmt.Println("nothing to show")
	}

	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

/*
Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
*/

func (l *LinkedList) ReturnKthFromLast(ind int) *Node {
	if l.Head == nil || ind < 0 {
		return nil
	}

	p1, p2 := l.Head, l.Head
	for range ind {
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
	}
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	return p1
}
