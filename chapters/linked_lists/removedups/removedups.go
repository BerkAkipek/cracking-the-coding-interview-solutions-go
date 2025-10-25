package removedups

import (
	"fmt"
	"reflect"
)

/*
Remove Dups: Write code to remove duplicates from an unsorted linked list.
Follow Up:
How would you solve this problem if a temporary buffer is not allowed?
*/
type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Append(val int) {
	newNode := &Node{Value: val}
	if l.Size == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *LinkedList) Prepend(val int) {
	newNode := &Node{Value: val}
	if l.Size == 0 {
		l.Head = newNode
		l.Tail = newNode
		return
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Size++
}

func (l *LinkedList) Insert(index, val int) (bool, error) {
	newNode := &Node{Value: val}
	if index < 0 || index > l.Size {
		return false, fmt.Errorf("index out of range")
	}

	if index == 0 {
		l.Prepend(val)
		return true, nil
	}

	if index == l.Size-1 {
		l.Append(val)
		return true, nil
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	l.Size++
	return true, nil
}

func (l *LinkedList) DeleteNode(val int) {
	if l.Head == nil {
		return
	}

	if val == l.Head.Value {
		l.Head = l.Head.Next
		l.Size--
		return
	}

	if val == l.Tail.Value {
		if reflect.DeepEqual(l.Head, l.Tail) {
			l.Head = nil
			l.Tail = nil
			l.Size--
			return
		}
		current := l.Head
		for current.Next != l.Tail {
			current = current.Next
		}
		l.Tail = current
		l.Tail.Next = nil
		l.Size--
		return
	}

	current, prev := l.Head.Next, l.Head
	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
			l.Size--
			return
		}
		prev = current
		current = current.Next
	}
}

func (l *LinkedList) PrintList() {
	if l.Size == 0 {
		fmt.Println("no Elements")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (l *LinkedList) RemoveDups() LinkedList {
	seen := make(map[int]bool)
	result := &LinkedList{}

	current := l.Head
	for current != nil {
		if !seen[current.Value] {
			seen[current.Value] = true
			result.Append(current.Value)
		}
		current = current.Next
	}
	return *result
}
