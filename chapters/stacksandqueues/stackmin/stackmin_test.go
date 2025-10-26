package stackmin

import "testing"

func TestPushAndMin(t *testing.T) {
	tests := []struct {
		name     string
		sequence []int
		expected []int // expected mins after each push
	}{
		{
			name:     "decreasing sequence",
			sequence: []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "increasing sequence",
			sequence: []int{1, 2, 3, 4, 5},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "mixed sequence",
			sequence: []int{5, 6, 3, 7, 2},
			expected: []int{5, 5, 3, 3, 2},
		},
		{
			name:     "with duplicates",
			sequence: []int{2, 2, 1, 1},
			expected: []int{2, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Stack
			for i, val := range tt.sequence {
				s.Push(val)
				min, _ := s.Min()
				if min != tt.expected[i] {
					t.Errorf("after Push(%d), got min=%d, want %d", val, min, tt.expected[i])
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	var s Stack
	sequence := []int{5, 6, 3, 7, 2}
	for _, val := range sequence {
		s.Push(val)
	}

	expectedPops := []struct {
		wantValue int
		wantMin   int
	}{
		{2, 3},
		{7, 3},
		{3, 5},
		{6, 5},
		{5, 0}, // stack empty after last pop
	}

	for i, exp := range expectedPops {
		got, ok := s.Pop()
		if !ok {
			t.Fatalf("unexpected empty stack at iteration %d", i)
		}
		if got != exp.wantValue {
			t.Errorf("Pop() got=%d, want=%d", got, exp.wantValue)
		}
		if min, ok := s.Min(); ok {
			if min != exp.wantMin {
				t.Errorf("after Pop(%d), got min=%d, want=%d", got, min, exp.wantMin)
			}
		}
	}
}

func TestEdgeCases(t *testing.T) {
	var s Stack

	// Pop from empty stack
	if _, ok := s.Pop(); ok {
		t.Error("expected Pop() to fail on empty stack")
	}

	// Min from empty stack
	if _, ok := s.Min(); ok {
		t.Error("expected Min() to fail on empty stack")
	}

	// Push duplicate mins
	s.Push(3)
	s.Push(3)
	s.Push(3)
	if min, _ := s.Min(); min != 3 {
		t.Errorf("expected min=3, got %d", min)
	}

	s.Pop()
	s.Pop()
	s.Pop()
	if _, ok := s.Min(); ok {
		t.Error("expected empty stack after popping all elements")
	}
}
