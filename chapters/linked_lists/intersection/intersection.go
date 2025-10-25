package intersection

/*
Intersection: Given two (singly) linked lists, determine if the two lists intersect.
Return the intersecting node. Note that the intersection is defined based on reference, not value.
That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second
linked list, then they are intersecting.
*/

type Node struct {
	Value int
	Next  *Node
}

type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist) Length() int {
	if l.Head == nil {
		return 0
	}
	count, current := 0, l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (l *Linkedlist) GetTail() *Node {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

func (l *Linkedlist) AppendList(nd *Node) {
	if l.Head == nil {
		l.Head = nd
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = nd
}

func IsIntersection(l1, l2 *Linkedlist) *Node {
	if l1.Head == nil || l2.Head == nil {
		return nil
	}

	if l1.GetTail() != l2.GetTail() {
		return nil
	}

	len1 := l1.Length()
	len2 := l2.Length()
	current1, current2 := l1.Head, l2.Head
	for len1 > len2 {
		len1--
		current1 = current1.Next
	}
	for len2 > len1 {
		len2--
		current2 = current2.Next
	}

	for current1 != nil && current2 != nil {
		if current1 == current2 {
			return current1
		}
		current1 = current1.Next
		current2 = current2.Next
	}

	return nil
}
