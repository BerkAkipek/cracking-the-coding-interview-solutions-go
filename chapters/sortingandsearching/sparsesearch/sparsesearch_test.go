package sparsesearch

import "testing"

func TestSparseSearch(t *testing.T) {
	tests := []struct {
		name   string
		arr    []string
		target string
		want   []int
	}{
		{
			name:   "basic example from prompt",
			arr:    []string{"at", "", "", "", "ball", "", "car", "", "", "dad", "", ""},
			target: "ball",
			want:   []int{4},
		},
		{
			name:   "target at start",
			arr:    []string{"ant", "", "", "bat", "", "cat"},
			target: "ant",
			want:   []int{0},
		},
		{
			name:   "target at end",
			arr:    []string{"ant", "", "bat", "", "", "dog"},
			target: "dog",
			want:   []int{5},
		},
		{
			name:   "target in middle with empties",
			arr:    []string{"ant", "", "", "", "bat", "", "cat", "", "", "dog"},
			target: "cat",
			want:   []int{6},
		},
		{
			name:   "target not found",
			arr:    []string{"ant", "", "bat", "", "cat"},
			target: "dog",
			want:   []int{-1},
		},
		{
			name:   "all empty strings",
			arr:    []string{"", "", "", "", ""},
			target: "ball",
			want:   []int{-1},
		},
		{
			name:   "empty array",
			arr:    []string{},
			target: "ball",
			want:   []int{-1},
		},
		{
			name:   "empty target",
			arr:    []string{"a", "b", "c"},
			target: "",
			want:   []int{-1},
		},
		{
			name:   "duplicate targets (any valid index acceptable)",
			arr:    []string{"ant", "", "bat", "bat", "", "", "bat"},
			target: "bat",
			want:   []int{2, 3, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SparseSearch(tt.arr, tt.target)
			valid := false
			for _, w := range tt.want {
				if got == w {
					valid = true
					break
				}
			}
			if !valid {
				t.Errorf("SparseSearch(%v, %q) = %d, want one of %v", tt.arr, tt.target, got, tt.want)
			}
		})
	}
}
