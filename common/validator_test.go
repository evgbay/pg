package common

import (
	"reflect"
	"strings"
	"testing"
)

func TestPasswordValidatorInstanse(t *testing.T) {
	t.Run("Successful create PasswordValidator struct", func(t *testing.T) {
		expected := &PasswordValidator{
			MinLength: 5,
			MaxLength: 10,
		}
		actual := NewPasswordValidator(5, 10)

		if actual == nil {
			t.Fatal("expected non-nil PasswordValidator")
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %+v, got %+v", expected, actual)
		}
	})

	t.Run("Should return empty PasswordValidator struct when input values are not correct", func(t *testing.T) {
		expected := &PasswordValidator{}
		actual := NewPasswordValidator(-1, 10)

		if actual == nil {
			t.Fatal("expected non-nil PasswordValidator")
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %+v, got %+v", expected, actual)
		}
	})
}

func TestPasswordValidatorValidateMethod(t *testing.T) {
	validator := &PasswordValidator{
		MinLength: 5,
		MaxLength: 10,
	}

	t.Run("Successful password length validation", func(t *testing.T) {
		input := 8
		err := validator.Validate(&input)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Error while password length is smaller than minLength value", func(t *testing.T) {
		input := 3
		err := validator.Validate(&input)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		if !strings.Contains(err.Error(), "Minimum password length - 5") {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Error while password length is greater than maxLength value", func(t *testing.T) {
		input := 11
		err := validator.Validate(&input)

		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		if !strings.Contains(err.Error(), "Maximum password length - 10") {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
