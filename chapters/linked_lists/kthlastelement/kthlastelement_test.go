package kthlastelement

import (
	"reflect"
	"testing"
)

func buildList(vals ...int) *LinkedList {
	list := &LinkedList{}
	for _, v := range vals {
		list.Append(v)
	}
	return list
}

func toSlice(l *LinkedList) []int {
	var out []int
	for n := l.Head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

func intPtr(v int) *int {
	return &v
}

func TestLinkedListLogic(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *LinkedList
		action   func(l *LinkedList)
		expected []int
	}{
		{
			name: "Append maintains correct order",
			setup: func() *LinkedList {
				l := &LinkedList{}
				l.Append(10)
				l.Append(20)
				l.Append(30)
				return l
			},
			action:   func(l *LinkedList) {},
			expected: []int{10, 20, 30},
		},
		{
			name: "Remove deletes head correctly",
			setup: func() *LinkedList {
				return buildList(10, 20, 30)
			},
			action: func(l *LinkedList) {
				l.Remove(10)
			},
			expected: []int{20, 30},
		},
		{
			name: "Remove deletes tail correctly",
			setup: func() *LinkedList {
				return buildList(10, 20, 30)
			},
			action: func(l *LinkedList) {
				l.Remove(30)
			},
			expected: []int{10, 20},
		},
		{
			name: "Remove ignores nonexistent value",
			setup: func() *LinkedList {
				return buildList(1, 2, 3)
			},
			action: func(l *LinkedList) {
				l.Remove(100)
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "Remove handles duplicates correctly",
			setup: func() *LinkedList {
				return buildList(10, 20, 20, 30)
			},
			action: func(l *LinkedList) {
				l.Remove(20)
			},
			expected: []int{10, 30},
		},
		{
			name: "Remove from empty list does nothing",
			setup: func() *LinkedList {
				return &LinkedList{}
			},
			action: func(l *LinkedList) {
				l.Remove(10)
			},
			expected: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := tc.setup()
			tc.action(list)
			got := toSlice(list)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("%s failed: got=%v; expected=%v", tc.name, got, tc.expected)
			}
		})
	}
}

func TestReturnKthFromLast(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *LinkedList
		k        int
		expected *int // use pointer to handle nil cases
	}{
		{
			name: "Last element (k=1)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30, 40, 50)
			},
			k:        1,
			expected: intPtr(50),
		},
		{
			name: "Second to last element (k=2)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30, 40, 50)
			},
			k:        2,
			expected: intPtr(40),
		},
		{
			name: "Middle element (k=3)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30, 40, 50)
			},
			k:        3,
			expected: intPtr(30),
		},
		{
			name: "First element (k=5)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30, 40, 50)
			},
			k:        5,
			expected: intPtr(10),
		},
		{
			name: "Too large (k > list length)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30)
			},
			k:        4,
			expected: nil,
		},
		{
			name: "Zero index (k=0)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30)
			},
			k:        0,
			expected: nil,
		},
		{
			name: "Negative index (k=-1)",
			setup: func() *LinkedList {
				return buildList(10, 20, 30)
			},
			k:        -1,
			expected: nil,
		},
		{
			name: "Empty list (k=1)",
			setup: func() *LinkedList {
				return &LinkedList{}
			},
			k:        1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.setup()
			got := list.ReturnKthFromLast(tt.k)

			switch {
			case tt.expected == nil && got != nil:
				t.Errorf("Expected nil, got %v", got.Value)
			case tt.expected != nil && got == nil:
				t.Errorf("Expected %v, got nil", *tt.expected)
			case tt.expected != nil && got.Value != *tt.expected:
				t.Errorf("Expected %v, got %v", *tt.expected, got.Value)
			}
		})
	}
}
