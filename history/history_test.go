package history

import "testing"

func TestLoad(t *testing.T) {
	t.Run("history load", func(t *testing.T) {
		if _, err := Load(); err != nil {
			t.Fatal(err)
		}
	})
}
