package queueviastacks

import (
	"reflect"
	"testing"
)

// Assume Stack has Push, Pop, Peek, Length, IsEmpty methods.

func TestQueueStacks(t *testing.T) {
	tests := []struct {
		name     string
		actions  func(q *QueueStacks) ([]int, []bool)
		expected []int
	}{
		{
			name: "enqueue and dequeue basic FIFO order",
			actions: func(q *QueueStacks) ([]int, []bool) {
				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)
				results, oks := []int{}, []bool{}
				for i := 0; i < 3; i++ {
					val, ok := q.Dequeue()
					results = append(results, val)
					oks = append(oks, ok)
				}
				return results, oks
			},
			expected: []int{1, 2, 3},
		},
		{
			name: "dequeue from empty queue should return false",
			actions: func(q *QueueStacks) ([]int, []bool) {
				val, ok := q.Dequeue()
				return []int{val}, []bool{ok}
			},
			expected: []int{0}, // value irrelevant, ok=false expected
		},
		{
			name: "interleaved enqueue and dequeue operations",
			actions: func(q *QueueStacks) ([]int, []bool) {
				q.Enqueue(10)
				q.Enqueue(20)
				v1, _ := q.Dequeue() // should give 10
				q.Enqueue(30)
				q.Enqueue(40)
				v2, _ := q.Dequeue() // should give 20
				q.Enqueue(50)
				v3, _ := q.Dequeue() // should give 30
				v4, _ := q.Dequeue() // should give 40
				v5, _ := q.Dequeue() // should give 50
				return []int{v1, v2, v3, v4, v5}, nil
			},
			expected: []int{10, 20, 30, 40, 50},
		},
		{
			name: "multiple dequeue calls with lazy transfer",
			actions: func(q *QueueStacks) ([]int, []bool) {
				for i := 1; i <= 5; i++ {
					q.Enqueue(i)
				}
				vals := []int{}
				for i := 0; i < 3; i++ {
					v, _ := q.Dequeue()
					vals = append(vals, v)
				}
				q.Enqueue(6)
				q.Enqueue(7)
				for i := 0; i < 4; i++ {
					v, _ := q.Dequeue()
					vals = append(vals, v)
				}
				return vals, nil
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "peek should not remove elements",
			actions: func(q *QueueStacks) ([]int, []bool) {
				q.Enqueue(100)
				q.Enqueue(200)
				v1, _ := q.Peek() // should see 100
				v2, _ := q.Dequeue()
				v3, _ := q.Peek() // should see 200
				v4, _ := q.Dequeue()
				return []int{v1, v2, v3, v4}, nil
			},
			expected: []int{100, 100, 200, 200},
		},
		{
			name: "check IsEmpty works correctly",
			actions: func(q *QueueStacks) ([]int, []bool) {
				results := []int{}
				if !q.IsEmpty() {
					t.Errorf("expected empty queue at start")
				}
				q.Enqueue(5)
				if q.IsEmpty() {
					t.Errorf("queue should not be empty after enqueue")
				}
				q.Dequeue()
				if !q.IsEmpty() {
					t.Errorf("queue should be empty after removing all elements")
				}
				return results, nil
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueueStacks()
			results, _ := tt.actions(q)
			if len(tt.expected) > 0 && !reflect.DeepEqual(results, tt.expected) {
				t.Errorf("%s: got %v, want %v", tt.name, results, tt.expected)
			}
		})
	}
}
