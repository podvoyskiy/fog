package config

import "testing"

func TestLoad(t *testing.T) {
	t.Run("load config", func(t *testing.T) {
		tmpDir := t.TempDir()
		if _, err := Load(tmpDir); err != nil {
			t.Fatal(err)
		}
	})
}
