package threinone

import "fmt"

/*
Three in One: Describe how you could use a single array to implement three stacks.
*/

type MultiStack struct {
	values []int
	size   [3]int
	cap    int
}

func NewMultiStack(stackCapacity int) *MultiStack {
	return &MultiStack{
		values: make([]int, stackCapacity*3),
		cap:    stackCapacity,
	}
}

func (ms *MultiStack) indexOfTop(stackNum int) int {
	offset := stackNum * ms.cap
	return offset + ms.size[stackNum] - 1
}

func (ms *MultiStack) Push(stackNum, value int) error {
	if ms.size[stackNum] >= ms.cap {
		return fmt.Errorf("stack %d is full", stackNum)
	}
	ms.size[stackNum]++
	ms.values[ms.indexOfTop(stackNum)] = value
	return nil
}

func (ms *MultiStack) Pop(stackNum int) (int, error) {
	if ms.size[stackNum] == 0 {
		return 0, fmt.Errorf("empty")
	}
	top := ms.indexOfTop(stackNum)
	val := ms.values[top]
	ms.values[top] = 0
	ms.size[stackNum]--
	return val, nil
}
