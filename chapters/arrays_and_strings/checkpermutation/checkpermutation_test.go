package checkpermutation

import "testing"

func TestCheckPermutation(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected bool
	}{
		// Basic positive cases
		{"Simple permutation", "abc", "bca", true},
		{"Reversed string", "abcd", "dcba", true},
		{"Same string", "aabb", "aabb", true},

		// Negative cases (different length or character set)
		{"Different lengths", "abc", "ab", false},
		{"Different characters", "abc", "abd", false},
		{"Extra repeated char", "aabb", "ababab", false},
		{"One empty, one not", "", "a", false},

		// Edge cases
		{"Both empty", "", "", true},
		{"Single char same", "x", "x", true},
		{"Single char different", "x", "y", false},
		{"Permutation with space between - true", "a b", "b a", true},
		{"Permutation with trailing space", "a b", "ab ", true},

		// Case sensitivity
		{"Different cases", "Abc", "abc", false},

		// Unicode and special characters
		{"Unicode permutation", "😀😁", "😁😀", true},
		{"Unicode not permutation", "😀😁", "😀😀", false},
		{"Symbols permutation", "!@#", "#@!", true},
		{"Symbols not permutation", "!@#", "#@@", false},

		// Complex repeated patterns
		{"Long repeated true", "aabbcc", "abcabc", true},
		{"Long repeated false", "aabbcc", "aabbc", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CheckPermutation(tc.str1, tc.str2)
			if result != tc.expected {
				t.Errorf("CheckPermutation(%q, %q) = %v; want %v",
					tc.str1, tc.str2, result, tc.expected)
			}
		})
	}
}
