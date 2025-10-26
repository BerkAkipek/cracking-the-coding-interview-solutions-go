package stackofplates

import (
	"reflect"
	"testing"
)

func stacksToSlices(s *StackOfStacks) [][]int {
	result := [][]int{}
	for _, st := range s.stacks {
		result = append(result, append([]int(nil), st.data...))
	}
	return result
}

func TestStackOfStacks(t *testing.T) {
	tests := []struct {
		name           string
		capacity       int
		ops            func(s *StackOfStacks) []int // sequence of operations, returns pops
		expectedStacks [][]int                      // final state
		expectedPops   []int
	}{
		{
			name:     "Single push and pop",
			capacity: 3,
			ops: func(s *StackOfStacks) []int {
				s.Push(10)
				val, _ := s.Pop()
				return []int{val}
			},
			expectedStacks: [][]int{},
			expectedPops:   []int{10},
		},
		{
			name:     "Push fills one stack and creates new one",
			capacity: 2,
			ops: func(s *StackOfStacks) []int {
				for i := 1; i <= 5; i++ {
					s.Push(i)
				}
				return nil
			},
			expectedStacks: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name:     "Pop removes from last stack correctly",
			capacity: 2,
			ops: func(s *StackOfStacks) []int {
				for i := 1; i <= 5; i++ {
					s.Push(i)
				}
				val, _ := s.Pop() // should remove 5
				return []int{val}
			},
			expectedStacks: [][]int{{1, 2}, {3, 4}},
			expectedPops:   []int{5},
		},
		{
			name:     "PopAt removes from middle stack",
			capacity: 2,
			ops: func(s *StackOfStacks) []int {
				for i := 1; i <= 6; i++ {
					s.Push(i)
				}
				val, _ := s.PopAt(1) // pop from middle stack [3,4]
				return []int{val}
			},
			expectedStacks: [][]int{{1, 2}, {3}, {5, 6}},
			expectedPops:   []int{4},
		},
		{
			name:     "PopAt removes empty stack and shifts correctly",
			capacity: 1,
			ops: func(s *StackOfStacks) []int {
				for i := 1; i <= 3; i++ {
					s.Push(i)
				}
				val, _ := s.PopAt(1) // remove [2]
				return []int{val}
			},
			expectedStacks: [][]int{{1}, {3}},
			expectedPops:   []int{2},
		},
		{
			name:     "PopAt last stack decrements currentStack correctly",
			capacity: 2,
			ops: func(s *StackOfStacks) []int {
				for i := 1; i <= 4; i++ {
					s.Push(i)
				}
				val, _ := s.PopAt(1) // pop from last stack
				return []int{val}
			},
			expectedStacks: [][]int{{1, 2}, {3}},
			expectedPops:   []int{4},
		},
		{
			name:     "Pop on empty structure returns false",
			capacity: 3,
			ops: func(s *StackOfStacks) []int {
				_, ok := s.Pop()
				if ok {
					t.Errorf("Expected false on Pop from empty structure")
				}
				return nil
			},
			expectedStacks: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStackOfStacks(tt.capacity)
			popped := tt.ops(s)

			gotStacks := stacksToSlices(s)
			if !reflect.DeepEqual(gotStacks, tt.expectedStacks) {
				t.Errorf("%s: got stacks = %v, want %v", tt.name, gotStacks, tt.expectedStacks)
			}

			if len(tt.expectedPops) > 0 && !reflect.DeepEqual(popped, tt.expectedPops) {
				t.Errorf("%s: got pops = %v, want %v", tt.name, popped, tt.expectedPops)
			}
		})
	}
}
