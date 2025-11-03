package sortedmerge

import (
	"reflect"
	"testing"
)

func TestSortedMerge(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "simple merge",
			a:    []int{1, 3, 5, 0, 0, 0},
			b:    []int{2, 4, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "B smaller elements",
			a:    []int{5, 6, 7, 0, 0, 0},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3, 5, 6, 7},
		},
		{
			name: "B larger elements",
			a:    []int{1, 2, 3, 0, 0, 0},
			b:    []int{4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "A has all smaller elements and B empty",
			a:    []int{1, 2, 3},
			b:    []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "A empty buffer, B non-empty",
			a:    []int{0, 0, 0},
			b:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Duplicate values across A and B",
			a:    []int{1, 3, 5, 0, 0, 0},
			b:    []int{3, 5, 7},
			want: []int{1, 3, 3, 5, 5, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortedMerge(append([]int{}, tt.a...), tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedMerge(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
