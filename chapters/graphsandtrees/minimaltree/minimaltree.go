package minimaltree

/*
Minimal Tree: Given a sorted (increasing order) array with unique integer elements,
write an algorithm to create a binary search tree with minimal height.
*/

type Node struct {
	val   int
	left  *Node
	right *Node
}

func build(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start+end) / 2
	root := &Node{val: arr[mid]}
	root.left = build(arr, start, mid-1)
	root.right = build(arr, mid+1, end)
	return root
}

func FromArrayToTree(arr []int) *Node {
	return build(arr, 0, len(arr)-1)
}
