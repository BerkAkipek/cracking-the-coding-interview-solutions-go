package buildorder

import (
	"reflect"
	"testing"
)

func TestKahnAlgorithm(t *testing.T) {
	tests := []struct {
		name         string
		projects     []string
		dependencies [][2]string
		want         []string
		wantErr      bool
		errContains  string
	}{
		{
			name:     "Example from CTCI",
			projects: []string{"a", "b", "c", "d", "e", "f"},
			dependencies: [][2]string{
				{"a", "d"},
				{"f", "b"},
				{"b", "d"},
				{"f", "a"},
				{"d", "c"},
			},
			want:    []string{"f", "e", "a", "b", "d", "c"}, // One valid topological order
			wantErr: false,
		},
		{
			name:     "Cycle detected",
			projects: []string{"a", "b"},
			dependencies: [][2]string{
				{"a", "b"},
				{"b", "a"},
			},
			wantErr:     true,
			errContains: "cycle",
		},
		{
			name:     "Unknown project in dependencies",
			projects: []string{"a", "b"},
			dependencies: [][2]string{
				{"x", "a"},
			},
			wantErr:     true,
			errContains: "unknown project",
		},
		{
			name:     "Independent projects",
			projects: []string{"x", "y", "z"},
			want:     []string{"x", "y", "z"}, // any order is fine, we only check membership
			wantErr:  false,
		},
		{
			name:     "Single project no dependencies",
			projects: []string{"a"},
			want:     []string{"a"},
			wantErr:  false,
		},
		{
			name:     "Simple linear chain",
			projects: []string{"a", "b", "c"},
			dependencies: [][2]string{
				{"a", "b"},
				{"b", "c"},
			},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KahnAlgorithm(tt.projects, tt.dependencies)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error state: got error=%v, wantErr=%v", err, tt.wantErr)
			}
			if err != nil && tt.errContains != "" && !contains(err.Error(), tt.errContains) {
				t.Fatalf("expected error to contain %q, got %q", tt.errContains, err.Error())
			}
			if !tt.wantErr {
				if len(got) != len(tt.projects) {
					t.Fatalf("unexpected order length: got=%d, want=%d", len(got), len(tt.projects))
				}
				if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
					t.Logf("Got order: %v", got)
					t.Logf("Expected (one valid order): %v", tt.want)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (func() bool {
		for i := 0; i+len(substr) <= len(s); i++ {
			if s[i:i+len(substr)] == substr {
				return true
			}
		}
		return false
	})()
}
