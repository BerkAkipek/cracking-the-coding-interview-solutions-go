package partition

import "fmt"

/*
Partition: Write code to partition a linked list around a value x, such that all nodes less than x come
before all nodes greater than or equal to x. If x is contained within the list, the values of x only need
to be after the elements less than x (see below). The partition element x can appear anywhere in the
"right partition"; it does not need to appear between the left and right partitions.
EXAMPLE
Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition= 5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
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

func (l *LinkedList) Prepend(val int) {
	newNode := &Node{Value: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList) Insert(ind, val int) {
	newNode := &Node{Value: val}
	if ind < 0 {
		return
	}

	if l.Head == nil && ind == 0 {
		l.Head = newNode
		return
	}

	if l.Head != nil && ind == 0 {
		l.Prepend(newNode.Value)
		return
	}

	i := 0
	prev, current := l.Head, l.Head
	for ; current != nil && i < ind; i++ {
		prev = current
		current = current.Next
	}
	if i < ind && current == nil {
		return
	}
	prev.Next = newNode
	newNode.Next = current
}

func (l *LinkedList) Delete(val int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == val {
		l.Head = l.Head.Next
		return
	}

	var prev *Node = nil
	current := l.Head
	for current != nil {
		if current.Value == val {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
}

func (l *LinkedList) DeleteAt(ind int) {
	if l.Head == nil {
		return
	}

	if ind < 0 {
		return
	}

	if l.Head != nil && ind == 0 {
		l.Head = l.Head.Next
		return
	}

	i := 1
	prev, current := l.Head, l.Head.Next
	for ; current != nil && i < ind; i++ {
		prev = current
		current = current.Next
	}
	if i < ind && current == nil {
		return
	}
	prev.Next = current.Next
}

func (l *LinkedList) ToSlice() []int {
	result := []int{}
	if l.Head == nil {
		return result
	}

	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

func (l *LinkedList) PrintAll() {
	if l.Head == nil {
		fmt.Println("nil")
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
}

func Partition(l *LinkedList, pivot int) *LinkedList {
	left, right := &LinkedList{}, &LinkedList{}
	if l.Head == nil {
		return left
	}
	current := l.Head
	for current != nil {
		if current.Value < pivot {
			left.Append(current.Value)
		} else {
			right.Append(current.Value)
		}
		current = current.Next
	}

	if left.Head == nil {
		return right
	}

	current = left.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = right.Head
	return left
}
