package common

import "testing"

func TestNewTemplate(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		special  *bool
		expected Template
	}{
		{
			name:    "Should create template with special symbols",
			length:  10,
			special: boolPtr(true),
			expected: Template{
				Len:     10,
				Letters: true,
				Digits:  true,
				Special: true,
			},
		},
		{
			name:    "Should create template without special symbols",
			length:  8,
			special: boolPtr(false),
			expected: Template{
				Len:     8,
				Letters: true,
				Digits:  true,
				Special: false,
			},
		},
		{
			name:    "Should create template with nil special symbols",
			length:  12,
			special: nil,
			expected: Template{
				Len:     12,
				Letters: true,
				Digits:  true,
				Special: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewTemplate(&tt.length, tt.special)

			if actual == nil {
				t.Fatal("expected non-nil Template")
			}

			if *actual != tt.expected {
				t.Errorf("expected %+v, got %+v", tt.expected, *actual)
			}
		})
	}
}

func boolPtr(value bool) *bool {
	return &value
}