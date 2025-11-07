package smallestdifference

import "testing"

func TestSmallestDifference(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want int
	}{
		{
			name: "Example from book",
			arr1: []int{1, 3, 15, 11, 2},
			arr2: []int{23, 127, 235, 19, 8},
			want: 3, // (11, 8)
		},
		{
			name: "Arrays with negatives",
			arr1: []int{-10, -5, 0, 5, 10},
			arr2: []int{-6, -2, 1, 7, 12},
			want: 1, // (-5, -6) or (0,1)
		},
		{
			name: "Overlapping ranges",
			arr1: []int{5, 10, 15},
			arr2: []int{10, 20, 30},
			want: 0, // (10,10)
		},
		{
			name: "Single element arrays",
			arr1: []int{5},
			arr2: []int{9},
			want: 4, // (5,9)
		},
		{
			name: "Identical arrays",
			arr1: []int{1, 2, 3},
			arr2: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "Empty first array",
			arr1: []int{},
			arr2: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "Empty second array",
			arr1: []int{1, 2, 3},
			arr2: []int{},
			want: 0,
		},
		{
			name: "Arrays with duplicates",
			arr1: []int{1, 1, 1, 5, 6},
			arr2: []int{2, 2, 2, 7, 8},
			want: 1,
		},
		{
			name: "Large difference between arrays",
			arr1: []int{1, 2, 3},
			arr2: []int{100, 200, 300},
			want: 97,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SmallestDifference(tt.arr1, tt.arr2)
			if got != tt.want {
				t.Errorf("SmallestDifference(%v, %v) = %d; want %d", tt.arr1, tt.arr2, got, tt.want)
			}
		})
	}
}
