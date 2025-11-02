package checkbalanced

/*
Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of
this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any
node never differ by more than one.
*/

type Node struct {
	left  *Node
	right *Node
}

func BalanceDFS(root *Node) bool {
	result := returnHeight(root)
	return result != -1
}

func returnHeight(n *Node) int {
	if n == nil {
		return 0
	}
	leftHeight := returnHeight(n.left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := returnHeight(n.right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	return 1 + max(leftHeight, rightHeight)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
