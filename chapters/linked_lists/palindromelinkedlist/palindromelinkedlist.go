package palindromelinkedlist

import (
	"fmt"
)

/*
Palindrome: Implement a function to check if a linked list is a palindrome

Input: 1 → 2 → 3 → 2 → 1 -> true
Input: 1 → 2 → 3 → 4 -> false
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
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Size++
}

func (l *LinkedList) Prepend(val int) {
	newNode := &Node{Value: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
	l.Size++
}

func BuildFromSlice(slc []int) *LinkedList {
	res := &LinkedList{}
	if len(slc) == 0 {
		return res
	}

	for _, elem := range slc {
		res.Append(elem)
	}
	return res
}

func (l *LinkedList) ToSlice() []int {
	slc := []int{}
	if l.Head == nil {
		return slc
	}

	current := l.Head
	for current != nil {
		slc = append(slc, current.Value)
		current = current.Next
	}
	return slc
}

func (l *LinkedList) PrintList() {
	if l.Head == nil {
		fmt.Printf("nil")
		return
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Printf("nil\n")
}

func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList) Length() int {
	return l.Size
}

/*
Palindrome() — using fast/slow pointer and reverse-half technique.
*/
func (l *LinkedList) Reverse() {
	if l.Head == nil {
		return
	}

	var track *Node
	var prev *Node
	current := l.Head
	for current != nil {
		track = current.Next
		current.Next = prev
		prev = current
		current = track
	}
	l.Head = prev
}

func (l *LinkedList) Remove(n int) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == n {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		l.Size--
		return
	}

	prev, current := l.Head, l.Head.Next
	for current != nil {
		if current.Value == n {
			prev.Next = current.Next
			l.Size--
			if current == l.Tail {
				l.Tail = prev
			}
			return
		}
		prev = current
		current = current.Next
	}
}

func RemoveDups(l *LinkedList) {
	seen := make(map[int]bool)
	if l.Head == nil {
		return
	}

	seen[l.Head.Value] = true
	prev, current := l.Head, l.Head.Next
	for current != nil {
		if !seen[current.Value] {
			seen[current.Value] = true
		} else {
			prev.Next = current.Next
			continue
		}
		prev = current
		current = current.Next
	}
}

func (l *LinkedList) Partition(x int) *LinkedList {
	left, right := &LinkedList{}, &LinkedList{}
	if l.Head == nil {
		return left
	}

	current := l.Head
	for current != nil {
		if current.Value < x {
			left.Append(current.Value)
		} else {
			right.Append(current.Value)
		}
		current = current.Next
	}

	if left.Head == nil {
		return right
	}

	if right.Head == nil {
		return left
	}

	left.Tail.Next = right.Head
	left.Tail = right.Tail
	return left
}

func DeleteMiddle(node *Node) {
	if node == nil || node.Next == nil {
		return
	}

	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

func KthLastElement(ind int, l *LinkedList) *Node {
	if l.Head == nil || ind < 0 {
		return nil
	}

	p1, p2 := l.Head, l.Head
	for range ind {
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p1
}

func (l *LinkedList) IsPalindrome() bool {
	if l.Head == nil || l.Head.Next == nil {
		return true
	}

	// --- Step 1: Find the midpoint using slow/fast pointers ---
	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// --- Step 2: Reverse the second half ---
	var prev *Node
	current := slow
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// --- Step 3: Compare both halves ---
	left, right := l.Head, prev
	for right != nil {
		if left.Value != right.Value {
			return false
		}
		left = left.Next
		right = right.Next
	}

	// Optional Step 4: Restore the list (not required)
	return true
}
