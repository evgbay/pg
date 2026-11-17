package common

import (
	"slices"
	"testing"
)

func TestGetCharset(t *testing.T) {
	tests := []struct {
		name     string
		template Template
		expected []rune
	}{
		{
			name:     "Should return All when Special is true",
			template: Template{Special: true},
			expected: All,
		},
		{
			name:     "Should return Alpha when default Template",
			template: Template{},
			expected: Alpha,
		},
		{
			name: "Should prioritize Special over other flags",
			template: Template{
				Special: true,
				Letters: true,
				Digits:  true,
			},
			expected: All,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetCharset(&tt.template)

			if !slices.Equal(tt.expected, actual) {
				t.Errorf("expected %q, got %q",
					string(tt.expected),
					string(actual))
			}
		})
	}
}
