package oneaway

import "testing"

func TestOneAway(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		target   string
		expected bool
	}{
		{"Remove char", "pale", "ple", true},
		{"Insert char", "pales", "pale", true},
		{"Replace char", "pale", "bale", true},
		{"Two edits", "pale", "bae", false},
		{"Identical", "pale", "pale", true},
		{"Completely different", "abc", "xyz", false},
		{"Empty and single", "", "a", true},
		{"Empty and double", "", "ab", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := OneAway(tc.input, tc.target)
			if result != tc.expected {
				t.Errorf("OneAway(%v) = %v; expected: %v", tc.input, result, tc.expected)
			}
		})
	}
}
