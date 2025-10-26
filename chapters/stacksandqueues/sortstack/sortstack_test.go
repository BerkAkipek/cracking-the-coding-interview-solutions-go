package sortstack

import (
	"reflect"
	"testing"
)

// helper to build a stack from slice (top = last element)
func buildStack(vals []int) *Stack {
	st := NewStack()
	for _, v := range vals {
		st.Push(v)
	}
	return st
}

// helper to convert stack to slice (top -> first)
func stackToSlice(s *Stack) []int {
	result := []int{}
	for !s.IsEmpty() {
		val, _ := s.Pop()
		result = append(result, val)
	}
	return result
}

func TestSortStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int // top -> bottom
	}{
		{
			name:     "empty stack",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "already sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted descending",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "contains duplicates",
			input:    []int{4, 1, 3, 4, 2, 1},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "random unsorted",
			input:    []int{3, 7, 2, 9, 1},
			expected: []int{1, 2, 3, 7, 9},
		},
		{
			name:     "large input check",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := buildStack(tt.input)
			sorted, ok := SortStack(s)
			if !ok {
				t.Fatalf("SortStack returned false for %s", tt.name)
			}
			got := stackToSlice(sorted)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s failed: got %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestStackOperations(t *testing.T) {
	st := NewStack()
	if !st.IsEmpty() {
		t.Fatal("new stack should be empty")
	}

	// Push/Pop
	st.Push(10)
	st.Push(20)
	st.Push(30)

	if v, _ := st.Peek(); v != 30 {
		t.Errorf("peek got %d, want 30", v)
	}

	v, ok := st.Pop()
	if !ok || v != 30 {
		t.Errorf("pop got (%d,%v), want (30,true)", v, ok)
	}

	v, ok = st.Pop()
	if !ok || v != 20 {
		t.Errorf("pop got (%d,%v), want (20,true)", v, ok)
	}

	v, ok = st.Pop()
	if !ok || v != 10 {
		t.Errorf("pop got (%d,%v), want (10,true)", v, ok)
	}

	// empty again
	if !st.IsEmpty() {
		t.Error("stack should be empty after popping all elements")
	}

	// pop from empty should fail gracefully
	if _, ok := st.Pop(); ok {
		t.Error("pop from empty stack should return ok=false")
	}
}
