package threeinone

/*
Three in One: Describe how you could use a single array to implement three stacks.
*/

type LinkedStack struct {
	Value int
	Next  int
}

type ThreeInOne struct {
	storage   []LinkedStack
	top       [3]int
	freeIndex int
}

func CreateMemory(size int) *ThreeInOne {
	storage := make([]LinkedStack, size)
	for i := range size {
		storage[i].Next = i + 1
	}

	storage[size-1].Next = -1

	result := &ThreeInOne{
		storage:   storage,
		top:       [3]int{-1, -1, -1},
		freeIndex: 0,
	}
	return result
}

/*
Push — handles empty and overflow conditions cleanly.

Pop — safely returns elements and recycles freed nodes.

Peek — lets you read the top without mutating.

IsEmpty — simple utility to improve code clarity.
*/

func (t *ThreeInOne) Push(stackNum, val int) (int, bool) {
	if stackNum < 0 || stackNum > 2 {
		return 1, false
	}

	if t.freeIndex == -1 {
		return 1, false
	}
	t.freeIndex = t.storage[t.freeIndex].Next
	t.storage[t.freeIndex].Value = val
	t.storage[t.freeIndex].Next = t.top[stackNum]
	t.top[stackNum] = t.freeIndex
	return 0, true
}
