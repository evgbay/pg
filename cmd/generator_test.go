package cmd

import (
	"github.com/evgbay/pg/common"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name     string
		template common.Template
	}{
		{
			name: "Should generate password with letters and digits",
			template: common.Template{
				Len:     10,
				Letters: true,
				Digits:  true,
				Special: false,
			},
		},
		{
			name: "Should generate password with special symbols",
			template: common.Template{
				Len:     16,
				Letters: true,
				Digits:  true,
				Special: true,
			},
		},
		{
			name: "Should generate password with letters only",
			template: common.Template{
				Len:     8,
				Letters: true,
				Digits:  false,
				Special: false,
			},
		},
		{
			name: "Should generate empty password when length is zero",
			template: common.Template{
				Len:     0,
				Letters: true,
				Digits:  true,
				Special: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GeneratePassword(&tt.template)
			charset := common.GetCharset(&tt.template)

			if len([]rune(actual)) != tt.template.Len {
				t.Errorf(
					"expected password length %d, got %d",
					tt.template.Len,
					len([]rune(actual)),
				)
			}

			for _, char := range actual {
				if !strings.ContainsRune(string(charset), char) {
					t.Errorf(
						"password contains invalid character %q",
						char,
					)
				}
			}
		})
	}
}
