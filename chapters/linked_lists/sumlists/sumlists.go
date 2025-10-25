package sumlists

import "fmt"

/*
Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
The digits are stored in reverse order, such that the 1 's digit is at the head of the list.
Write a function that adds the two numbers and returns the sum as a linked list.
EXAMPLE
Input: (7-> 1 -> 6) + (5 -> 9 -> 2).That is,617 + 295.
Output: 2 -> 1 -> 9. That is, 912.
FOLLOW UP
Suppose the digits are stored in forward order. Repeat the above problem.
EXAMPLE
lnput:(6 -> 1 -> 7) + (2 -> 9 -> 5).That is,617 + 295.
Output: 9 -> 1 -> 2. That is, 912.

In imperative programming, I tell the machine what to do.
In functional programming, I describe what something is, and let time reveal it.
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

func (l *LinkedList) AppendList(val int) {
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

func (l *LinkedList) PrintList() {
	if l.Head == nil {
		fmt.Printf("nil")
	}

	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
}

// ReverseList reverse given linked list
// It had written for satify the follow up question
func (l *LinkedList) ReverseList() {
	if l.Head == nil {
		return
	}

	var prev *Node
	var track *Node
	var current = l.Head

	for current != nil {
		track = current.Next
		current.Next = prev
		prev = current
		current = track
	}
	l.Head = prev
}

func BuildFromSlice(slc []int) *LinkedList {
	list := &LinkedList{}

	if len(slc) == 0 {
		return list
	}

	for _, elem := range slc {
		list.AppendList(elem)
	}
	return list
}

func helper(dgt1, dgt2, carry int, res *LinkedList) int {
	sum := dgt1 + dgt2 + carry
	step := sum % 10
	carry = sum / 10
	res.AppendList(step)
	return carry
}

func CarryTraversalSummation(l1, l2 *LinkedList) (*LinkedList, error) {
	result := &LinkedList{}
	if l1.Head == nil && l2.Head == nil {
		return result, nil
	}
	if l1.Head == nil {
		return BuildFromSlice(l2.ToSlice()), nil
	}
	if l2.Head == nil {
		return BuildFromSlice(l1.ToSlice()), nil
	}

	h1, h2 := l1.Head, l2.Head
	carry := 0
	for h1 != nil && h2 != nil {
		carry = helper(h1.Value, h2.Value, carry, result)
		h1 = h1.Next
		h2 = h2.Next
	}
	if h1 == nil && h2 == nil {
		if carry != 0 {
			result.AppendList(carry)
		}
		return result, nil
	} else {
		remaining := (*Node)(nil)
		if h1 != nil {
			remaining = h1
		} else {
			remaining = h2
		}
		for remaining != nil {
			carry = helper(remaining.Value, 0, carry, result)
			remaining = remaining.Next
		}
		if carry != 0 {
			result.AppendList(carry)
		}
	}
	return result, nil
}
