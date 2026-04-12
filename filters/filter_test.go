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
	pattern := "ls"
	t.Logf("pattern: %q", pattern)

	for _, typeF := range GetFilterTypes() {
		f, err := FromUint8(typeF.uint8())
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("=== filter: %s ===", typeF.toString())

		matches := f.Match(commands, pattern)

		for i, m := range matches {
			t.Logf("%d: score=%d, cmd=%q", i, m.Score, commands[m.Index])
		}
	}
}
