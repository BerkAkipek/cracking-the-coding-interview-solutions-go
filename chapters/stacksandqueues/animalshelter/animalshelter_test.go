package animalshelter

import (
	"fmt"
	"testing"
)

// setup helper: creates a shelter and enqueues animals
func setupShelter(t *testing.T, animals ...Animal) *Shelter {
	t.Helper()
	s := NewShelter()
	for _, a := range animals {
		s.Enqueue(a)
	}
	return s
}

func TestMakeNoise(t *testing.T) {
	tests := []struct {
		name     string
		animal   Animal
		expected string
	}{
		{"Dog barks", &Dog{}, "Bark"},
		{"Cat miavs", &Cat{}, "Miav"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.animal.MakeNoise()
			if got != tt.expected {
				t.Errorf("MakeNoise() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEnqueueAndDequeueAny(t *testing.T) {
	d1, _ := NewDog()
	c1, _ := NewCat()
	d2, _ := NewDog()

	s := setupShelter(t, d1, c1, d2)

	tests := []struct {
		name         string
		action       func() (Animal, error)
		wantType     string
		expectError  bool
		remainingLen int
	}{
		{
			name:         "Dequeue first (Dog)",
			action:       func() (Animal, error) { return s.DequeueAny() },
			wantType:     fmt.Sprintf("%T", &Dog{}),
			expectError:  false,
			remainingLen: 2,
		},
		{
			name:         "Dequeue second (Cat)",
			action:       func() (Animal, error) { return s.DequeueAny() },
			wantType:     fmt.Sprintf("%T", &Cat{}),
			expectError:  false,
			remainingLen: 1,
		},
		{
			name:         "Dequeue third (Dog)",
			action:       func() (Animal, error) { return s.DequeueAny() },
			wantType:     fmt.Sprintf("%T", &Dog{}),
			expectError:  false,
			remainingLen: 0,
		},
		{
			name:         "Dequeue from empty",
			action:       func() (Animal, error) { return s.DequeueAny() },
			wantType:     "",
			expectError:  true,
			remainingLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := tt.action()
			if (err != nil) != tt.expectError {
				t.Fatalf("expected error=%v, got %v", tt.expectError, err)
			}
			if !tt.expectError && a == nil {
				t.Fatalf("expected animal, got nil")
			}
			if a != nil && tt.wantType != "" {
				gotType := fmt.Sprintf("%T", a)
				if gotType != tt.wantType {
					t.Errorf("expected type %v, got %v", tt.wantType, gotType)
				}
			}
			if len(s.adoptionQueue) != tt.remainingLen {
				t.Errorf("adoptionQueue len = %d, want %d", len(s.adoptionQueue), tt.remainingLen)
			}
		})
	}
}

func TestDequeueDogAndCat(t *testing.T) {
	d1, _ := NewDog()
	d2, _ := NewDog()
	c1, _ := NewCat()

	s := setupShelter(t, d1, d2, c1)

	tests := []struct {
		name        string
		action      func() (interface{}, error)
		wantType    string
		expectError bool
	}{
		{
			name:        "DequeueDog returns Dog",
			action:      func() (interface{}, error) { return s.DequeueDog() },
			wantType:    fmt.Sprintf("%T", &Dog{}),
			expectError: false,
		},
		{
			name:        "DequeueCat returns Cat",
			action:      func() (interface{}, error) { return s.DequeueCat() },
			wantType:    fmt.Sprintf("%T", &Cat{}),
			expectError: false,
		},
		{
			name:        "DequeueDog again",
			action:      func() (interface{}, error) { return s.DequeueDog() },
			wantType:    fmt.Sprintf("%T", &Dog{}),
			expectError: false,
		},
		{
			name:        "DequeueDog from empty",
			action:      func() (interface{}, error) { return s.DequeueDog() },
			wantType:    "",
			expectError: true,
		},
		{
			name:        "DequeueCat from empty",
			action:      func() (interface{}, error) { return s.DequeueCat() },
			wantType:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.action()
			if (err != nil) != tt.expectError {
				t.Fatalf("expected error=%v, got %v", tt.expectError, err)
			}
			if !tt.expectError && got == nil {
				t.Fatalf("expected non-nil, got nil")
			}
			if got != nil && tt.wantType != "" {
				gotType := fmt.Sprintf("%T", got)
				if gotType != tt.wantType {
					t.Errorf("expected type %v, got %v", tt.wantType, gotType)
				}
			}
		})
	}
}
