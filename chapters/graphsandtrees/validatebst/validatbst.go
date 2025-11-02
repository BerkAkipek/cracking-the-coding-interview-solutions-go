package validatebst

/*
Validate BST: Implement a function to check if a binary tree is a binary search tree.
*/
type Node struct {
	val   int
	left  *Node
	right *Node
}

func isValid(root *Node, minimum, maximum int) bool {
	if root == nil {
		return true
	}

	if root.val <= minimum || root.val >= maximum {
		return false
	}

	return isValid(root.left, minimum, root.val) && isValid(root.right, root.val, maximum)
}

func IsValid(root *Node) bool {
	inf := (1 << 63) - 1
	minInf := -1 << 63
	return isValid(root, minInf, inf)
}
