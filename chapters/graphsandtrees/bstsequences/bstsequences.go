package bstsequences

/*
BST Sequences: A binary search tree was created by traversing through an array from left to right
and inserting each element. Given a binary search tree with distinct elements, print all possible
arrays that could have led to this tree.
EXAMPLE
Output: {2, 1, 3}, {2, 3, 1}
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (n *Node) Insert(val int) *Node {
	if n == nil {
		n = &Node{
			val:   val,
			left:  nil,
			right: nil,
		}
		return n
	}

	if val < n.val {
		n.left = n.left.Insert(val)
	} else {
		n.right = n.right.Insert(val)
	}

	return n
}

func (n *Node) BuildOrder() [][]int {
	if n == nil {
		return [][]int{{}}
	}

	leftSeq := n.left.BuildOrder()
	rightSeq := n.right.BuildOrder()
	prefix := []int{n.val}
	result := [][]int{}

	for _, left := range leftSeq {
		for _, right := range rightSeq {
			var weaved [][]int
			weaveEngine(left, right, prefix, &weaved)
			result = append(result, weaved...)
		}
	}
	return result
}

func weaveEngine(left, right, prefix []int, result *[][]int) {
	if len(left) == 0 || len(right) == 0 {
		out := append(append([]int{}, prefix...), append(left, right...)...)
		*result = append(*result, out)
		return
	}
	weaveEngine(left[1:], right, append(prefix, left[0]), result)
	weaveEngine(left, right[1:], append(prefix, right[0]), result)
}
