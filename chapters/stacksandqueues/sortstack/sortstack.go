package sortstack

/*
Sort Stack: Write a program to sort a stack such that the smallest items are on the top.
You can use an additional temporary stack, but you may not copy the elements into any other data structure
(such as an array). The stack supports the following operations: push, pop, peek, and is Empty.
*/

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: []int{},
	}
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	element := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return element, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	return s.data[len(s.data)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func SortStack(s *Stack) (*Stack, bool) {
	if s.IsEmpty() {
		return s, true
	}

	tempStorage := NewStack()
	for !s.IsEmpty() {
		temp, ok := s.Pop()
		if !ok {
			return nil, false
		}

		if tempStorage.IsEmpty() {
			tempStorage.Push(temp)
			continue
		}

		lookUp, ok := tempStorage.Peek()
		if !ok {
			break
		}
		for lookUp > temp {
			step, ok := tempStorage.Pop()
			if !ok {
				return nil, false
			}
			s.Push(step)
			lookUp, ok = tempStorage.Peek()
			if !ok {
				break
			}
		}
		tempStorage.Push(temp)
	}
	for !tempStorage.IsEmpty() {
		val, _ := tempStorage.Pop()
		s.Push(val)
	}
	return s, true
}
