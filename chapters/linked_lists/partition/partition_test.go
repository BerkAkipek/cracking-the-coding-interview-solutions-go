package partition

import (
	"reflect"
	"testing"
)

func TestLinkedListLogic(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*LinkedList)
		expected []int
	}{
		// --- Append ---
		{
			name: "Append to empty list",
			setup: func(l *LinkedList) {
				l.Append(10)
			},
			expected: []int{10},
		},
		{
			name: "Append multiple elements",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
			},
			expected: []int{10, 20, 30},
		},

		// --- Prepend ---
		{
			name: "Prepend to empty list",
			setup: func(l *LinkedList) {
				l.Prepend(10)
			},
			expected: []int{10},
		},
		{
			name: "Prepend to non-empty list",
			setup: func(l *LinkedList) {
				l.Append(20)
				l.Prepend(10)
			},
			expected: []int{10, 20},
		},

		// --- Insert ---
		{
			name: "Insert at head",
			setup: func(l *LinkedList) {
				l.Append(20)
				l.Insert(0, 10)
			},
			expected: []int{10, 20},
		},
		{
			name: "Insert in middle",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(30)
				l.Insert(1, 20)
			},
			expected: []int{10, 20, 30},
		},
		{
			name: "Insert at out of range index (ignored)",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Insert(5, 20)
			},
			expected: []int{10},
		},

		// --- Delete by value ---
		{
			name: "Delete head node",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Delete(10)
			},
			expected: []int{20},
		},
		{
			name: "Delete middle node",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.Delete(20)
			},
			expected: []int{10, 30},
		},
		{
			name: "Delete non-existing value (ignored)",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Delete(50)
			},
			expected: []int{10, 20},
		},

		// --- DeleteAt ---
		{
			name: "DeleteAt head index",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.DeleteAt(0)
			},
			expected: []int{20},
		},
		{
			name: "DeleteAt middle index",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.DeleteAt(1)
			},
			expected: []int{10, 30},
		},
		{
			name: "DeleteAt out of bounds (ignored)",
			setup: func(l *LinkedList) {
				l.Append(10)
				l.DeleteAt(10)
			},
			expected: []int{10},
		},

		// --- Edge cases ---
		{
			name: "Operations on empty list should not panic",
			setup: func(l *LinkedList) {
				l.Delete(10)
				l.DeleteAt(0)
				l.Insert(2, 50)
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			tt.setup(l)
			got := l.ToSlice()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s failed: got %v, expected %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		pivot    int
		expected []int
	}{
		{
			name:     "Mixed nodes around pivot",
			input:    []int{3, 5, 8, 5, 10, 2, 1},
			pivot:    5,
			expected: []int{3, 2, 1, 5, 8, 5, 10},
		},
		{
			name:     "All less than pivot",
			input:    []int{1, 2, 3},
			pivot:    5,
			expected: []int{1, 2, 3},
		},
		{
			name:     "All greater or equal to pivot",
			input:    []int{8, 9, 10},
			pivot:    5,
			expected: []int{8, 9, 10},
		},
		{
			name:     "Empty list",
			input:    []int{},
			pivot:    5,
			expected: []int{},
		},
		{
			name:     "All equal to pivot",
			input:    []int{5, 5, 5},
			pivot:    5,
			expected: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}
			for _, v := range tt.input {
				l.Append(v)
			}

			result := Partition(l, tt.pivot)
			got := result.ToSlice()

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Partition(%v, %v) = %v; expected %v",
					tt.input, tt.pivot, got, tt.expected)
			}
		})
	}
}
