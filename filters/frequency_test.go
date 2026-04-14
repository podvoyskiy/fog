package filters

import (
	"fmt"
	"testing"
)

func TestFrequencyFilter(t *testing.T) {
	f, err := FromUint8(typeFrequency.uint8())
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		commands []string
		pattern  string
		want     []string
	}{
		{
			commands: []string{"cd ~", "alias | grep 'cd'", "cd", "cd ~", "ls"},
			pattern:  "cd",
			want:     []string{"cd ~", "cd", "alias | grep 'cd'"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d_%s", i, tt.pattern), func(t *testing.T) {
			matches := f.Match(tt.commands, tt.pattern)

			if len(matches) != len(tt.want) {
				t.Fatalf("got %d matches | want %d matches", len(matches), len(tt.want))
			}

			for i, match := range matches {
				cmd := tt.commands[match.Index]
				if cmd != tt.want[i] {
					t.Errorf("position %d: got %q | want %q", i, cmd, tt.want[i])
				}
			}
		})
	}
}
