package jenny

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Assert on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("Assert on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "Grace", "hello")
	})
}
