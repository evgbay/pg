package common

import (
	"testing"
)

func TestToValueMethod(t *testing.T) {
	t.Run("Should return value from string pointer", func(t *testing.T) {
		str := "OK"
		val := ToValue(&str)

		if val != str {
			t.Errorf("expected %q, got %q", str, val)
		}
	})

	t.Run("Should return value from int pointer", func(t *testing.T) {
		num := 5
		val := ToValue(&num)

		if val != num {
			t.Errorf("expected %d, got %d", num, val)
		}
	})

	t.Run("Should return default type value if input pointer nil", func(t *testing.T) {
		var ptr *int
		actual := ToValue(ptr)

		if actual != 0 {
			t.Errorf("expected 0, got %d", actual)
		}
	})
}
