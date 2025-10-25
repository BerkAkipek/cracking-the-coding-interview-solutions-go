package palindromelinkedlist

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected bool
	}{
		// --- Basic cases ---
		{"Empty list", []int{}, true},                  // by definition: empty = palindrome
		{"Single element", []int{1}, true},             // single node always palindrome
		{"Two same elements", []int{5, 5}, true},       // symmetric pair
		{"Two different elements", []int{5, 7}, false}, // not symmetric

		// --- Odd-length palindromes ---
		{"Odd palindrome", []int{1, 2, 3, 2, 1}, true},
		{"Odd non-palindrome", []int{1, 2, 3, 4, 1}, false},

		// --- Even-length palindromes ---
		{"Even palindrome", []int{1, 2, 2, 1}, true},
		{"Even non-palindrome", []int{1, 2, 3, 4}, false},

		// --- Larger and asymmetric patterns ---
		{"Long palindrome", []int{1, 2, 3, 4, 3, 2, 1}, true},
		{"Almost palindrome", []int{1, 2, 3, 2, 2}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			l := BuildFromSlice(tc.values)
			got := l.IsPalindrome()
			if got != tc.expected {
				t.Errorf("IsPalindrome(%v) = %v; want %v", tc.values, got, tc.expected)
			}
		})
	}
}
