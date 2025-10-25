package deletemiddlenode

import (
	"reflect"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		deleteAt int
		expected []int
	}{
		{"delete middle node", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"delete node near end", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"delete head node", []int{1, 2, 3}, 0, []int{2, 3}},
		{"delete tail node", []int{1, 2, 3}, 2, []int{1, 2, 3}},
		{"single element list", []int{1}, 0, []int{1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &LinkedList{}
			for _, v := range tc.values {
				list.Append(v)
			}
			deletedElement := list.NodeAt(tc.deleteAt)
			DeleteMiddleNode(deletedElement)
			slc := list.ToSlice()
			if !reflect.DeepEqual(slc, tc.expected) {
				t.Errorf("expected: %v; got: %v", tc.expected, slc)
			}
		})
	}
}
