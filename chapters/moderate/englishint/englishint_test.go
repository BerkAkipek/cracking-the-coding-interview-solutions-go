package englishint

import "testing"

func TestEnglishInt(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		// Basic single digits
		{"Zero", 0, "zero"},
		{"One", 1, "one"},
		{"Five", 5, "five"},
		{"Nine", 9, "nine"},

		// Teens
		{"Ten", 10, "ten"},
		{"Eleven", 11, "eleven"},
		{"Nineteen", 19, "nineteen"},

		// Tens
		{"Twenty", 20, "twenty"},
		{"FortyTwo", 42, "forty two"},
		{"NinetyNine", 99, "ninety nine"},

		// Hundreds
		{"OneHundred", 100, "one hundred"},
		{"OneHundredFive", 105, "one hundred five"},
		{"ThreeHundredFortyTwo", 342, "three hundred forty two"},

		// Thousands
		{"OneThousand", 1000, "one thousand"},
		{"OneThousandTwoHundredThirtyFour", 1234, "one thousand two hundred thirty four"},
		{"FortyFiveThousand", 45000, "forty five thousand"},
		{"SeventyThousandEightHundredNinety", 70890, "seventy thousand eight hundred ninety"},

		// Millions and beyond
		{"OneMillion", 1000000, "one million"},
		{"TwoMillionThreeHundredFortyFiveThousandSixHundredSeventyEight",
			2345678, "two million three hundred forty five thousand six hundred seventy eight"},
		{"OneBillion", 1000000000, "one billion"},

		// Quadrillions and quintillions (upper safe limits)
		{"OneTrillion", 1000000000000, "one trillion"},
		{"OneQuadrillion", 1000000000000000, "one quadrillion"},
		{"OneQuintillion", 1000000000000000000, "one quintillion"},

		// Negative numbers
		{"NegativeOne", -1, "negative one"},
		{"NegativeFortyTwo", -42, "negative forty two"},
		{"NegativeMillion", -1000000, "negative one million"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EnglishInt(tt.input)
			if got != tt.expected {
				t.Errorf("EnglishInt(%d) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
