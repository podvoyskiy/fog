package filters

import "testing"

func TestFilterDebug(t *testing.T) {
	commands := []string{
		"ls",
		"ls -la",
		"alias ls",
		"ls -la ./",
		"lsof -i",
		"lsblk",
		"lsblk -f",
		"lscpu",
		"cd ~/home/LS/tmp",
		"lsblk -f",
	}

	for _, pattern := range []string{"ls", "ls  "} {
		t.Logf("\npattern: %q", pattern)

		for _, typeF := range allFilterTypes() {
			f, err := FromUint8(typeF.uint8())
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("=== filter: %s ===", typeF)

			matches := f.Match(commands, pattern)

			for _, m := range matches {
				t.Logf("score=%d, cmd=%q", m.Score, commands[m.Index])
			}
		}
	}
}
