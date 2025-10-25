package stringrotation

import "testing"

func TestStringRotation(t *testing.T) {
	tests := []struct {
		name     string
		input1   string
		input2   string
		expected bool
	}{
		{"Basic rotation", "waterbottle", "erbottlewat", true},
		{"Same letters different order", "waterbottle", "ttlewaterbo", true},
		{"Identical strings", "rotation", "rotation", true},
		{"Different lengths", "abc", "ab", false},
		{"One empty string", "", "abc", false},
		{"Both empty strings", "", "", true},
		{"Small rotation", "ab", "ba", true},
		{"Partial rotation", "abcdef", "efabcdx", false},
		{"Repeating characters", "aaaa", "aaaa", true},
		{"Middle rotation", "rotationtest", "testrotation", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsRotation(tc.input1, tc.input2)
			if got != tc.expected {
				t.Errorf("IsRotation(%v, %v) = %v; expected: %v", tc.input1, tc.input2, got, tc.expected)
			}
		})
	}
}
