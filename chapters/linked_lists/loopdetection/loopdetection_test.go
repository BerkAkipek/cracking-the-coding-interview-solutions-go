package loopdetection

import (
	"testing"
)

// Helper to create loop at given position (-1 = no loop)
func createLoop(l *LinkedList, pos int) {
	if pos < 0 || l.Head == nil {
		return
	}

	loopNode := l.Head
	for i := 0; i < pos && loopNode.Next != nil; i++ {
		loopNode = loopNode.Next
	}

	tail := l.Head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = loopNode
}

func TestFLoydsAlgorithm(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		loopPos  int // -1 means no loop
		expected int // -1 means no loop expected
	}{
		{"Empty list", []int{}, -1, -1},
		{"Single node, no loop", []int{1}, -1, -1},
		{"Single node loops to itself", []int{1}, 0, 1},
		{"Multiple nodes, no loop", []int{1, 2, 3, 4, 5}, -1, -1},
		{"Loop starts at head", []int{1, 2, 3, 4, 5}, 0, 1},
		{"Loop starts in middle", []int{10, 20, 30, 40, 50}, 2, 30},
		{"Loop starts at tail (self-loop)", []int{5, 10, 15, 20}, 3, 20},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			list := &LinkedList{}
			for _, v := range tc.values {
				list.AppendList(v)
			}
			createLoop(list, tc.loopPos)

			result := FloydsAlgorithm(list)
			if tc.expected == -1 && result != nil {
				t.Errorf("Expected no loop, got loop start at %d", result.Value)
			}
			if tc.expected != -1 {
				if result == nil {
					t.Errorf("Expected loop start at %d, got nil", tc.expected)
				} else if result.Value != tc.expected {
					t.Errorf("Expected loop start at %d, got %d", tc.expected, result.Value)
				}
			}
		})
	}
}
