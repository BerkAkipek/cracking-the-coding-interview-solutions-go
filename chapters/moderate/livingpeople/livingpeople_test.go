package livingpeople

import "testing"

func TestMostPeopleAlive(t *testing.T) {
	tests := []struct {
		name     string
		people   []*People
		interval [2]int
		expected int
	}{
		{
			name: "single person full life",
			people: []*People{
				{Birth: 1900, Death: 1950},
			},
			interval: [2]int{1900, 2000},
			expected: 1900,
		},
		{
			name: "overlapping lifespans",
			people: []*People{
				{Birth: 1900, Death: 1950},
				{Birth: 1920, Death: 1980},
				{Birth: 1950, Death: 2000},
			},
			interval: [2]int{1900, 2000},
			expected: 1950,
		},
		{
			name: "non-overlapping lifespans",
			people: []*People{
				{Birth: 1900, Death: 1910},
				{Birth: 1920, Death: 1930},
				{Birth: 1940, Death: 1950},
			},
			interval: [2]int{1900, 2000},
			expected: 1900,
		},
		{
			name: "multiple same-year births and deaths",
			people: []*People{
				{Birth: 1900, Death: 1900},
				{Birth: 1900, Death: 1901},
				{Birth: 1901, Death: 1902},
			},
			interval: [2]int{1900, 1902},
			expected: 1900,
		},
		{
			name:     "empty list",
			people:   []*People{},
			interval: [2]int{1900, 2000},
			expected: 1900,
		},
		{
			name: "narrow interval subset",
			people: []*People{
				{Birth: 1980, Death: 1990},
				{Birth: 1990, Death: 1995},
				{Birth: 1992, Death: 1999},
			},
			interval: [2]int{1985, 1995},
			expected: 1990,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MostPeopleAlive(tt.people, tt.interval)
			if got != tt.expected {
				t.Errorf("MostPeopleAlive(%v, %v) = %d; want %d", tt.people, tt.interval, got, tt.expected)
			}
		})
	}
}
