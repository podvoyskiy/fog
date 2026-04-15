package filters

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	u "github.com/podvoyskiy/fog/utils"
)

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

		for _, typeF := range []FilterType{typeDefault, typeFuzzy, typeFrequency} {
			f := newFilter(typeF)
			t.Logf("=== filter: %s ===", typeF)

			matches := f.Match(commands, pattern)

			for _, m := range matches {
				t.Logf("score=%d, cmd=%q", m.Score, commands[m.Index])
			}
		}
	}
}

func TestFilterRace(t *testing.T) {
	commands := generateFakeCommands(10000)
	runRace(t, commands, "ls")
}

func BenchmarkFilterRace(b *testing.B) {
	commands := generateFakeCommands(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runRace(b, commands, "ls")
	}
}

func runRace(t testing.TB, commands []string, pattern string) {
	start := time.Now()
	timeout := 100 * time.Millisecond

	type Result struct {
		typeF FilterType
		value int
		err   error
	}

	ch := make(chan Result, 2)

	filtering := func(typeF FilterType) {
		f := newFilter(typeF)

		matches := f.Match(commands, pattern)
		ch <- Result{typeF: typeF, value: len(matches)}
	}

	go filtering(typeFuzzy)
	go filtering(typeFrequency)

	select {
	case result := <-ch:
		if result.err != nil {
			u.Red().Printf("filter error: %v\n", result.err)
			t.Fatal()
		}
		msg := fmt.Sprintf("%s filter | length: %d | duration: %v", result.typeF, result.value, time.Since(start))
		if result.typeF == typeFrequency {
			u.Green().Println(msg)
		} else {
			u.Yellow().Println(msg)
		}
	case <-time.After(timeout):
		u.Red().Printf("timeout (%v) - no filter completed\n", timeout)
		t.Fatal()
	}
}

func generateFakeCommands(n int) []string {
	fakeCmds := make([]string, 0, n)

	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789 -_./=")

	for i := 0; i < n; i++ {
		length := rand.Intn(15) + 3
		cmd := make([]rune, length)
		for i := range cmd {
			cmd[i] = chars[rand.Intn(len(chars))]
		}
		fakeCmds = append(fakeCmds, string(cmd))
	}
	return fakeCmds
}
