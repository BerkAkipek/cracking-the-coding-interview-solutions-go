package divingboard

import (
	"reflect"
	"testing"
)

func TestDivingBoard(t *testing.T) {
	tests := []struct {
		name     string
		short    int
		long     int
		k        int
		expected []int
	}{
		{
			name:     "basic progression descending",
			short:    3,
			long:     5,
			k:        4,
			expected: []int{20, 18, 16, 14, 12},
		},
		{
			name:     "equal plank lengths",
			short:    4,
			long:     4,
			k:        3,
			expected: []int{12},
		},
		{
			name:     "k equals zero",
			short:    3,
			long:     5,
			k:        0,
			expected: nil,
		},
		{
			name:     "short greater than long",
			short:    6,
			long:     5,
			k:        3,
			expected: nil,
		},
		{
			name:     "one plank only",
			short:    2,
			long:     5,
			k:        1,
			expected: []int{5, 2},
		},
		{
			name:     "large k small difference",
			short:    5,
			long:     6,
			k:        5,
			expected: []int{30, 29, 28, 27, 26, 25},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DivingBoard(tt.short, tt.long, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("DivingBoard(%d, %d, %d) = %v; want %v",
					tt.short, tt.long, tt.k, got, tt.expected)
			}
		})
	}
}
