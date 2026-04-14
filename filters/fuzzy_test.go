package filters

import (
	"fmt"
	"testing"
)

func TestFuzzyFilter(t *testing.T) {
	f := &fuzzyFilter{}

	tests := []struct {
		commands []string
		pattern  string
		want     map[int]bool
	}{
		{
			commands: []string{"git commit", "gc", "commit git", "g c"},
			pattern:  "gc",
			want:     map[int]bool{0: true, 1: true, 3: true},
		},
		{
			commands: []string{"docker ps -a", "df -h", "docker images"},
			pattern:  "dc",
			want:     map[int]bool{0: true, 2: true},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d_%s", i, tt.pattern), func(t *testing.T) {
			matches := f.Match(tt.commands, tt.pattern)

			if len(matches) != len(tt.want) {
				t.Fatalf("got %d matches | want %d matches", len(matches), len(tt.want))
			}

			for _, match := range matches {
				if !tt.want[match.Index] {
					t.Errorf("unexpected index %d", match.Index)
				}
			}
		})
	}
}
