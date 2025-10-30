package firstcommonancestor

/*
First Common Ancestor: Design an algorithm and write code to find the first common ancestor
of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not
necessarily a binary search tree.
*/

type Node struct {
	left  *Node
	right *Node
}

func exists(root, target *Node) bool {
	if root == nil {
		return false
	}
	if root == target {
		return true
	}
	return exists(root.left, target) || exists(root.right, target)
}

func FirstCommonAncestor(root, x, y *Node) *Node {
	if !exists(root, x) || !exists(root, y) {
		return nil
	}
	return findAncestor(root, x, y)
}

func findAncestor(current, x, y *Node) *Node {
	if current == nil {
		return nil
	}

	if current == x || current == y {
		return current
	}

	left := findAncestor(current.left, x, y)
	right := findAncestor(current.right, x, y)

	if left != nil && right != nil {
		return current
	}

	if left != nil {
		return left
	}
	return right
}
