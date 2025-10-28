package checkbalanced

/*
Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of
this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any
node never differ by more than one
*/

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func size(n *TreeNode) int {
	if n == nil {
		return 0
	}

	leftSize := size(n.left)
	if leftSize == -1 {
		return -1
	}

	rightSize := size(n.right)
	if rightSize == -1 {
		return -1
	}

	if abs(leftSize-rightSize) > 1 {
		return -1
	}

	return max(leftSize, rightSize) + 1
}

func IsBalanced(n *TreeNode) bool {
	return size(n) != -1
}
