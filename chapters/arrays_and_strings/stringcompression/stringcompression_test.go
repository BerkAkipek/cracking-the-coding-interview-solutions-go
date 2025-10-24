package stringcompression

import "testing"

func TestStringCompression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// --- Basic cases ---
		{"Empty string", "", ""},
		{"Single character", "a", "a"},
		{"Two same characters", "aa", "aa"},      // same length → return original
		{"Two different characters", "ab", "ab"}, // compressed longer → return original

		// --- Typical compression ---
		{"Mixed run", "aabcccccaaa", "a2b1c5a3"},
		{"All same character", "aaaaaa", "a6"},
		{"No repeats", "abcdef", "abcdef"},
		{"Shorter original", "aabbcc", "aabbcc"},

		// --- Case sensitivity ---
		{"Upper and lower case", "AAaa", "AAaa"},
		{"Alternating cases", "AaAaAa", "AaAaAa"},

		// --- Edge patterns ---
		{"Single long run at start", "aaaaab", "a5b1"},
		{"Single long run at end", "baaaaa", "b1a5"},
		{"Repeating groups", "aaabbbccc", "a3b3c3"},
		{"One long run", "aaaaaaaaaaa", "a11"},

		// --- Unicode (sanity check for rune handling) ---
		{"Unicode letters", "äääöö", "ä3ö2"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := StringCompression(tc.input)
			if got != tc.expected {
				t.Errorf("StringCompression(%v) = %v; expected: %v", tc.input, got, tc.expected)
			}
		})
	}
}
