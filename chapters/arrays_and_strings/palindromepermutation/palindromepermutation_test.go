package palindromepermutation

import "testing"

func TestPalindromePermutation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Palindrome with spaces", "Tact Coa", true},
		{"Simple palindrome", "civic", true},
		{"Not a palindrome permutation", "hello", false},
		{"All same letters", "aaaa", true},
		{"Mixed case palindrome", "RaceCar", true},
		{"Non-letter characters", "A man, a plan, a canal, Panama!", true},
		{"Two odd counts", "abc", false},
		{"Empty string", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := PalindromePermutation(tc.input)
			if result != tc.expected {
				t.Errorf("PalindromePermutation(%v) = %v; expected: %v", tc.input, result, tc.expected)
			}
		})
	}
}
