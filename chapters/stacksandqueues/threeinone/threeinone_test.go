package threinone

import (
	"testing"
)

func TestMultiStack_PushPop(t *testing.T) {
	ms := NewMultiStack(3)

	if err := ms.Push(0, 10); err != nil {
		t.Fatalf("Push failed: %v", err)
	}
	if err := ms.Push(0, 20); err != nil {
		t.Fatalf("Push failed: %v", err)
	}

	_ = ms.Push(1, 100)
	_ = ms.Push(1, 200)

	_ = ms.Push(2, 999)

	v, err := ms.Pop(0)
	if err != nil {
		t.Fatalf("Pop failed: %v", err)
	}
	if v != 20 {
		t.Errorf("Pop() = %d; want 20", v)
	}

	v, err = ms.Pop(1)
	if err != nil {
		t.Fatalf("Pop failed: %v", err)
	}
	if v != 200 {
		t.Errorf("Pop() = %d; want 200", v)
	}

	v, err = ms.Pop(2)
	if err != nil {
		t.Fatalf("Pop failed: %v", err)
	}
	if v != 999 {
		t.Errorf("Pop() = %d; want 999", v)
	}

	v, err = ms.Pop(0)
	if err != nil {
		t.Fatalf("Pop failed: %v", err)
	}
	if v != 10 {
		t.Errorf("Pop() = %d; want 10", v)
	}
}

func TestMultiStack_Underflow(t *testing.T) {
	ms := NewMultiStack(2)

	_, err := ms.Pop(0)
	if err == nil {
		t.Errorf("Expected error on popping empty stack")
	}
}

func TestMultiStack_Overflow(t *testing.T) {
	ms := NewMultiStack(1)

	if err := ms.Push(0, 42); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	err := ms.Push(0, 99)
	if err == nil {
		t.Errorf("Expected overflow error, got nil")
	}
}

func TestMultiStack_IsolatedStacks(t *testing.T) {
	ms := NewMultiStack(2)

	ms.Push(0, 1)
	ms.Push(1, 10)
	ms.Push(2, 100)

	v0, _ := ms.Pop(0)
	v1, _ := ms.Pop(1)
	v2, _ := ms.Pop(2)

	if v0 != 1 || v1 != 10 || v2 != 100 {
		t.Errorf("Stacks are interfering: got (%d, %d, %d)", v0, v1, v2)
	}
}
