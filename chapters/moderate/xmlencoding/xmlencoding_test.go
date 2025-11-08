package xmlencoding

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		element  *Element
		mapping  map[string]int
		expected string
	}{
		{
			name: "basic family example",
			element: &Element{
				Tag: "family",
				Attributes: []Attribute{
					{"lastName", "McDowell"},
					{"state", "CA"},
				},
				Children: []*Element{
					{
						Tag: "person",
						Attributes: []Attribute{
							{"firstName", "Gayle"},
						},
						Value: "Some Message",
					},
				},
			},
			mapping: map[string]int{
				"family":    1,
				"person":    2,
				"firstName": 3,
				"lastName":  4,
				"state":     5,
			},
			expected: "1 4 McDowell 5 CA 0 2 3 Gayle 0 Some Message 0 0",
		},
		{
			name: "single element no attributes no children",
			element: &Element{
				Tag:   "person",
				Value: "Hello",
			},
			mapping: map[string]int{
				"person": 1,
			},
			expected: "1 0 Hello 0",
		},
		{
			name: "element with only attributes",
			element: &Element{
				Tag: "book",
				Attributes: []Attribute{
					{"title", "DSA"},
					{"author", "Skiena"},
				},
			},
			mapping: map[string]int{
				"book":   1,
				"title":  2,
				"author": 3,
			},
			expected: "1 2 DSA 3 Skiena 0 0",
		},
		{
			name: "nested multiple children",
			element: &Element{
				Tag: "root",
				Children: []*Element{
					{Tag: "child", Value: "A"},
					{Tag: "child", Value: "B"},
				},
			},
			mapping: map[string]int{
				"root":  1,
				"child": 2,
			},
			expected: "1 0 2 0 A 0 2 0 B 0 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.element, tt.mapping)
			if got != tt.expected {
				t.Errorf("Encode() = %q; want %q", got, tt.expected)
			}
		})
	}
}
