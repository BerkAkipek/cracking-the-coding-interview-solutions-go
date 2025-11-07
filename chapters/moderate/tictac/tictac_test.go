package tictac

import "testing"

func TestWhoWins(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]int
		expected int
	}{
		{
			name: "X wins first row",
			board: [][]int{
				{0, 0, 0},
				{1, -1, 1},
				{-1, 1, -1},
			},
			expected: 0,
		},
		{
			name: "O wins second row",
			board: [][]int{
				{0, -1, 0},
				{1, 1, 1},
				{-1, 0, -1},
			},
			expected: 1,
		},
		{
			name: "X wins first column",
			board: [][]int{
				{0, 1, -1},
				{0, 1, 1},
				{0, -1, -1},
			},
			expected: 0,
		},
		{
			name: "O wins last column",
			board: [][]int{
				{0, 1, 1},
				{-1, 0, 1},
				{-1, -1, 1},
			},
			expected: 1,
		},
		{
			name: "X wins main diagonal",
			board: [][]int{
				{0, 1, 1},
				{-1, 0, 1},
				{-1, -1, 0},
			},
			expected: 0,
		},
		{
			name: "O wins anti-diagonal",
			board: [][]int{
				{0, -1, 1},
				{0, 1, -1},
				{1, -1, 0},
			},
			expected: 1,
		},
		{
			name: "No winner yet",
			board: [][]int{
				{0, 1, 0},
				{1, 0, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "Empty board",
			board: [][]int{
				{-1, -1, -1},
				{-1, -1, -1},
				{-1, -1, -1},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WhoWins(tt.board)
			if got != tt.expected {
				t.Errorf("%s: WhoWins() = %d, want %d", tt.name, got, tt.expected)
			}
		})
	}
}
