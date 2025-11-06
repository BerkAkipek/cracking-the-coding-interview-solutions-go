package wordfrequencies

import (
	"reflect"
	"testing"
)

func TestWordFrequencies(t *testing.T) {
	tests := []struct {
		name string
		book []string
		want map[string]int
	}{
		{
			name: "basic single line",
			book: []string{"Hello world! Hello."},
			want: map[string]int{"hello": 2, "world": 1},
		},
		{
			name: "punctuation and casing",
			book: []string{"Book, book; BOOK!"},
			want: map[string]int{"book": 3},
		},
		{
			name: "ellipsis handling",
			book: []string{"Wait... wait… wait."},
			want: map[string]int{"wait": 3},
		},
		{
			name: "brackets and parentheses",
			book: []string{"(good) [bad] {ugly}"},
			want: map[string]int{"good": 1, "bad": 1, "ugly": 1},
		},
		{
			name: "mixed symbols around words",
			book: []string{"“hello!” …hello… --hello--"},
			want: map[string]int{"hello": 3},
		},
		{
			name: "empty and spaced lines",
			book: []string{"", "    ", "one two   two"},
			want: map[string]int{"one": 1, "two": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordFrequencies(tt.book)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WordFrequencies() = %v, want %v", got, tt.want)
			}
		})
	}
}
