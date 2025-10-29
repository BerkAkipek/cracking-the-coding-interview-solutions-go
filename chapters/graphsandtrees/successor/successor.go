package successor

/*
Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a binary search tree.
You may assume that each node has a link to its parent.
*/
type Node struct {
	left   *Node
	right  *Node
	parent *Node
}

func Successor(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.right != nil {
		next := n.right
		for next.left != nil {
			next = next.left
		}
		return next
	}

	current := n
	parent := current.parent
	for parent != nil && parent.right == current {
		current = parent
		parent = parent.parent
	}
	return parent
}
