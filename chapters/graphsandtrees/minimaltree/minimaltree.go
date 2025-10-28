package minimaltree

/*
Minimal Tree: Given a sorted (increasing order) array with unique integer elements,
write an algorithm to create a binary search tree with minimal height
*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

func CreateNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

type BinaryTree struct {
	Root *Node
}

func NewTree(start int) *BinaryTree {
	return &BinaryTree{
		Root: CreateNode(start),
	}
}

// BuildTree builds a balanced Sorted Binary Tree from a sorted array.
func BuildTree(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}
	middle := (start + end) / 2
	node := CreateNode(arr[middle])

	node.left = BuildTree(arr, start, middle-1)
	node.right = BuildTree(arr, middle+1, end)

	return node
}

func (t *BinaryTree) AddElement(val int) *Node {
	if t == nil {
		return nil
	}
	return insert(t.Root, val)
}

func insert(n *Node, val int) *Node {
	if n == nil {
		return CreateNode(val)
	}
	if val < n.data {
		n.left = insert(n.left, val)
		return n
	} else {
		n.right = insert(n.right, val)
		return n
	}
}
