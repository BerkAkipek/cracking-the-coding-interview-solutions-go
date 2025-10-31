package hundredlockers

import (
	"reflect"
	"testing"
)

// TestLockers uses table-driven tests to verify correctness.
func TestLockers(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want []int
	}{
		{
			name: "1 locker",
			num:  1,
			want: []int{1},
		},
		{
			name: "10 lockers",
			num:  10,
			want: []int{1, 4, 9},
		},
		{
			name: "20 lockers",
			num:  20,
			want: []int{1, 4, 9, 16},
		},
		{
			name: "50 lockers",
			num:  50,
			want: []int{1, 4, 9, 16, 25, 36, 49},
		},
		{
			name: "100 lockers",
			num:  100,
			want: []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Lockers(tc.num)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Lockers(%d) = %v; want %v", tc.num, got, tc.want)
			}
		})
	}
}
