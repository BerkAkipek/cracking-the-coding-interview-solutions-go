package urlify

import "testing"

func TestUrlify(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		trueLength int
		expected   string
	}{
		{
			name:       "Normal case",
			input:      "Mr John Smith    ",
			trueLength: 13,
			expected:   "Mr%20John%20Smith",
		},
		{
			name:       "Trailing buffer spaces",
			input:      "Hello World  ",
			trueLength: 11,
			expected:   "Hello%20World",
		},
		{
			name:       "Leading spaces",
			input:      "  Mr John Smith    ",
			trueLength: 15,
			expected:   "%20%20Mr%20John%20Smith",
		},
		{
			name:       "Multiple internal spaces",
			input:      "Go  Lang  Rocks      ",
			trueLength: 15,
			expected:   "Go%20%20Lang%20%20Rocks",
		},
		{
			name:       "No spaces",
			input:      "GoLang",
			trueLength: 6,
			expected:   "GoLang",
		},
		{
			name:       "All spaces",
			input:      "     ",
			trueLength: 5,
			expected:   "%20%20%20%20%20",
		},
		{
			name:       "Empty string",
			input:      "",
			trueLength: 0,
			expected:   "",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Urlify(testCase.input, testCase.trueLength)
			if testCase.expected != result {
				t.Errorf("Urliify(%v, %v) = %v; expected %v", testCase.input, testCase.trueLength, result, testCase.expected)
			}
		})
	}
}
