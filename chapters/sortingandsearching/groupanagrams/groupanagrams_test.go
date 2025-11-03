package groupanagrams

import (
	"slices"
	"testing"
)

func normalize(arr []string) []string {
	slices.Sort(arr)
	return arr
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string // groups of anagrams
	}{
		{
			name: "basic case",
			in:   []string{"bat", "tab", "eat", "tea", "tan", "nat"},
			want: [][]string{
				{"bat", "tab"},
				{"eat", "tea"},
				{"tan", "nat"},
			},
		},
		{
			name: "single word",
			in:   []string{"hello"},
			want: [][]string{{"hello"}},
		},
		{
			name: "mixed case",
			in:   []string{"abc", "bca", "xyz", "zyx", "foo"},
			want: [][]string{
				{"abc", "bca"},
				{"xyz", "zyx"},
				{"foo"},
			},
		},
		{
			name: "empty input",
			in:   []string{},
			want: [][]string{},
		},
		{
			name: "duplicate words",
			in:   []string{"tea", "eat", "tea"},
			want: [][]string{
				{"tea", "eat", "tea"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.in)

			normalize(got)

			var wantFlat []string
			for _, group := range tt.want {
				wantFlat = append(wantFlat, group...)
			}
			normalize(wantFlat)

			if !slices.Equal(got, wantFlat) {
				t.Errorf("GroupAnagrams(%v) = %v; want %v", tt.in, got, wantFlat)
			}
		})
	}
}
