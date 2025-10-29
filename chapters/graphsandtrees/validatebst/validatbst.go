package validatebst

/*
Validate BST: Implement a function to check if a binary tree is a binary search tree.
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

func IsValid(n *Node, start, stop int) bool {
	if n == nil {
		return true
	}

	if n.val <= start || n.val >= stop {
		return false
	}

	return IsValid(n.left, start, n.val) && IsValid(n.right, n.val, stop)
}
