package loopdetection

import (
	"fmt"
	"math/rand"
)

/*
Loop Detection: Given a circular linked list,
implement an algorithm that returns the node at the beginning of the loop.
DEFINITION
Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so
as to make a loop in the linked list.
EXAMPLE
Input: A -> B -> C -> D -> E -> C [the same C as earlier]
Output: C
*/

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) AppendList(val int) {
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

func (l *LinkedList) Display() {
	if l.Head == nil {
		fmt.Printf("nil\n")
	}
	current := l.Head
	for current.Next != nil {
		fmt.Printf("%v -> ", current.Value)
	}
	fmt.Printf("%v\n", current.Value)
}

func FromSlice(slc []int) *LinkedList {
	result := &LinkedList{}
	if len(slc) == 0 {
		return nil
	}
	for _, elem := range slc {
		result.AppendList(elem)
	}
	return result
}

func CreateLoop(l *LinkedList) {
	if l.Head == nil {
		return
	}

	interval := 0
	current := l.Head
	for current != nil {
		interval++
		current = current.Next
	}

	spice := rand.Intn(interval)
	interval = 0
	hold := l.Head
	current = l.Head
	for current.Next != nil {
		if interval == spice {
			hold = current
		}
		interval++
		current = current.Next
	}
	current.Next = hold
}

func FloydsAlgorithm(l *LinkedList) *Node {
	if l.Head == nil {
		return nil
	}

	slow, fast := l.Head, l.Head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = l.Head
			for {
				if fast == slow {
					return fast
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
}
