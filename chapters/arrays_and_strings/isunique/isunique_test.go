package isunique_test

import (
	"testing"

	"github.com/BerkAkipek/cracking-the-coding-interview-solutions-go/chapters/arrays_and_strings/isunique"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"All unique lowercase", "abcdef", true},
		{"Duplicate letters", "hello", false},
		{"Empty string", "", true},
		{"Single character", "a", true},
		{"Spaces repeated", "  ", false},
		{"Case-sensitive check", "GoLang", true}, // G != g
		{"Unicode unique", "🙂🙃😉", true},
		{"Unicode duplicate", "🙂🙃🙂", false},
		{"Digits unique", "12345", true},
		{"Digits repeated", "11234", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isunique.IsUnique(tt.input)
			if result != tt.expected {
				t.Errorf("IsUnique(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
