package removedups

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		inputs   func(*LinkedList)
		expected []int
	}{
		{
			name: "Append elements",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
			},
			expected: []int{10, 20, 30},
		},
		{
			name: "Prepend element",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Prepend(5)
			},
			expected: []int{5, 10, 20},
		},
		{
			name: "Insert in middle",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.Insert(1, 15)
			},
			expected: []int{10, 15, 20, 30},
		},
		{
			name: "Delete head node",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.DeleteNode(10)
			},
			expected: []int{20, 30},
		},
		{
			name: "Delete middle node",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.DeleteNode(20)
			},
			expected: []int{10, 30},
		},
		{
			name: "Delete tail node",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.DeleteNode(30)
			},
			expected: []int{10, 20},
		},
		{
			name: "Delete only element",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.DeleteNode(10)
			},
			expected: nil,
		},
		{
			name: "Delete non-existent value",
			inputs: func(l *LinkedList) {
				l.Append(10)
				l.Append(20)
				l.Append(30)
				l.DeleteNode(99)
			},
			expected: []int{10, 20, 30},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &LinkedList{}
			tc.inputs(list)
			current := list.Head
			var got []int
			for current != nil {
				got = append(got, current.Value)
				current = current.Next
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got: %v; expected: %v", got, tc.expected)
			}
		})
	}
}

func TestRemoveDups(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []int
		expected []int
	}{
		{
			name:     "All unique elements",
			inputs:   []int{10, 20, 30, 40},
			expected: []int{10, 20, 30, 40},
		},
		{
			name:     "All duplicates same value",
			inputs:   []int{5, 5, 5, 5, 5},
			expected: []int{5},
		},
		{
			name:     "Duplicates at start",
			inputs:   []int{10, 10, 20, 30},
			expected: []int{10, 20, 30},
		},
		{
			name:     "Duplicates at end",
			inputs:   []int{10, 20, 30, 30, 30},
			expected: []int{10, 20, 30},
		},
		{
			name:     "Alternating duplicates",
			inputs:   []int{10, 20, 10, 30, 20, 40},
			expected: []int{10, 20, 30, 40},
		},
		{
			name:     "Empty list",
			inputs:   []int{},
			expected: nil,
		},
		{
			name:     "Single element",
			inputs:   []int{42},
			expected: []int{42},
		},
		{
			name:     "Mixed duplicates",
			inputs:   []int{1, 2, 3, 2, 4, 1, 5, 3, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &LinkedList{}
			for _, v := range tc.inputs {
				list.Append(v)
			}

			result := list.RemoveDups()

			current := result.Head
			var got []int
			for current != nil {
				got = append(got, current.Value)
				current = current.Next
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("RemoveDups(%v) = %v; expected %v", tc.inputs, got, tc.expected)
			}
		})
	}
}
