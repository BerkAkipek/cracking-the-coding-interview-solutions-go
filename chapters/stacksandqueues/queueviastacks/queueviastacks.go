package queueviastacks

/*
Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks
*/

// LIFO
type Stack struct {
	data []int
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

func (s *Stack) Length() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

// FIFO
type QueueStacks struct {
	inStack  *Stack
	outStack *Stack
}

func NewQueueStacks() *QueueStacks {
	return &QueueStacks{
		inStack:  &Stack{},
		outStack: &Stack{},
	}
}

func (q *QueueStacks) Enqueue(val int) {
	q.inStack.Push(val)
}

func (q *QueueStacks) Dequeue() (int, bool) {
	if q.inStack.Length() == 0 && q.outStack.Length() == 0 {
		return 0, false
	}

	if q.outStack.Length() != 0 {
		element, ok := q.outStack.Pop()
		if !ok {
			return 0, false
		}
		return element, true
	}

	for q.inStack.Length() != 0 {
		element, ok := q.inStack.Pop()
		if !ok {
			return 0, false
		}
		q.outStack.Push(element)
	}

	element, ok := q.outStack.Pop()
	if !ok {
		return 0, false
	}
	return element, true
}

func (q *QueueStacks) IsEmpty() bool {
	return q.inStack.IsEmpty() && q.outStack.IsEmpty()
}

func (q *QueueStacks) Peek() (int, bool) {
	if q.inStack.Length() == 0 && q.outStack.Length() == 0 {
		return 0, false
	}

	if q.outStack.Length() != 0 {
		return q.outStack.data[q.outStack.Length()-1], true
	}

	for q.inStack.Length() != 0 {
		element, ok := q.inStack.Pop()
		if !ok {
			return 0, false
		}
		q.outStack.Push(element)
	}
	return q.outStack.data[q.outStack.Length()-1], true
}
