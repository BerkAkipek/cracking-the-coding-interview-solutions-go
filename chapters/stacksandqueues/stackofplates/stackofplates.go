package stackofplates

/*
Stack of Plates: Imagine a (literal) stack of plates.
If the stack gets too high, it might topple.
Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold.
Implement a data structure SetOfStacks that mimics this.
SetO-fStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity.
SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack
(that is, pop () should return the same values as it would if there were just a single stack).
FOLLOW UP
Implement a function popAt ( int index) which performs a pop operation on a specific sub-stack.
*/
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

func (s *Stack) Peek() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type StackOfStacks struct {
	stacks        []*Stack
	stackCapacity int
	currentStack  int
}

func NewStackOfStacks(capacity int) (*StackOfStacks, bool) {
	if capacity > 0 {
		return &StackOfStacks{
			stacks:        []*Stack{},
			stackCapacity: capacity,
			currentStack:  -1,
		}, true
	}
	return nil, false
}

func (s *StackOfStacks) Push(val int) {
	if s.currentStack == -1 || len(s.stacks[s.currentStack].data) == s.stackCapacity {
		newStack := &Stack{
			data: []int{},
		}
		newStack.Push(val)
		s.stacks = append(s.stacks, newStack)
		s.currentStack++
		return
	}

	s.stacks[s.currentStack].Push(val)
}

func (s *StackOfStacks) Pop() (int, bool) {
	if s.currentStack == -1 {
		return 0, false
	}

	if s.stacks[s.currentStack].Length() == 1 {
		currentElement, ok := s.stacks[s.currentStack].Pop()
		if !ok {
			return 0, false
		}
		s.stacks = s.stacks[:len(s.stacks)-1]
		s.currentStack--
		return currentElement, true
	}

	currentElement, ok := s.stacks[s.currentStack].Pop()
	if !ok {
		return 0, false
	}
	return currentElement, true
}

func (s *StackOfStacks) PopAt(index int) (int, bool) {
	if index < 0 || index > len(s.stacks)-1 {
		return 0, false
	}

	if s.stacks[index].Length() == 1 {
		currentElement, ok := s.stacks[index].Pop()
		if !ok {
			return 0, false
		}
		if s.stacks[index].Length() == 0 {
			s.stacks = append(s.stacks[:index], s.stacks[index+1:]...)
			s.currentStack--
		}
		return currentElement, true
	}

	currentElement, ok := s.stacks[index].Pop()
	if !ok {
		return 0, false
	}
	return currentElement, true
}
