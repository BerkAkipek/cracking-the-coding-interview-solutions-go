package sumlists

import (
	"reflect"
	"testing"
)

func TestCarryTraversalSummation(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{"Equal length no carry at end", []int{7, 1, 6}, []int{5, 9, 2}, []int{2, 1, 9}}, // 617 + 295 = 912
		{"First longer", []int{7, 1, 6}, []int{5, 9}, []int{2, 1, 7}},                    // 617 + 95 = 712
		{"Second longer", []int{5, 9}, []int{7, 1, 6}, []int{2, 1, 7}},                   // symmetric
		{"Carry extends length", []int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},            // 999 + 1 = 1000
		{"Both lists empty", []int{}, []int{}, []int{}},
		{"First empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"Second empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"Single digits no carry", []int{3}, []int{4}, []int{7}},                         // 3 + 4 = 7
		{"Single digits with carry", []int{5}, []int{8}, []int{3, 1}},                    // 5 + 8 = 13
		{"Multiple carries", []int{9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 1}}, // 9999 + 9999 = 19998
		{"All zeros", []int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}},                    // 0 + 0 = 0
		{"Different lengths with carry", []int{9, 9}, []int{1}, []int{0, 0, 1}},          // 99 + 1 = 100
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := BuildFromSlice(tt.l1)
			l2 := BuildFromSlice(tt.l2)

			got, _ := CarryTraversalSummation(l1, l2)
			gotSlice := got.ToSlice()

			if !reflect.DeepEqual(gotSlice, tt.expected) {
				t.Errorf("CarryTraversalSummation(%v, %v) = %v; want %v", tt.l1, tt.l2, gotSlice, tt.expected)
			}
		})
	}
}
