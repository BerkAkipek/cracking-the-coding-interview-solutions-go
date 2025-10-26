package stackmin

/*
Stack Min: How would you design a stack which, in addition to push and pop, has a function min
which returns the minimum element? Push, pop and min should all operate in 0(1) time
*/
type Stack struct {
	data []int
	min  []int
}

func (s *Stack) Push(val int) {
	if len(s.data) == 0 || len(s.min) == 0 {
		s.data = append(s.data, val)
		s.min = append(s.min, val)
		return
	}
	if val <= s.min[len(s.min)-1] {
		s.min = append(s.min, val)
	}
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}

	element := s.data[len(s.data)-1]

	if element == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}

	s.data = s.data[:len(s.data)-1]
	return element, true
}

func (s *Stack) Min() (int, bool) {
	if len(s.min) == 0 {
		return 0, false
	}
	return s.min[len(s.min)-1], true
}

func (s *Stack) Top() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1], true
}
