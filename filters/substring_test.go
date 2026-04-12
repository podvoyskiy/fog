package filters

import (
	"fmt"
	"testing"
)

func TestSubstringFilter(t *testing.T) {
	f, err := FromUint8(typeSubstring.uint8())
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		commands []string
		pattern  string
		want     []int
	}{
		{
			commands: []string{"foo", "bar", "baz"},
			pattern:  "foo",
			want:     []int{0},
		},
		{
			commands: []string{"foo", "bar", "baz"},
			pattern:  "bar",
			want:     []int{1},
		},
		{
			commands: []string{"Foo", "Bar", "Baz"},
			pattern:  "baz",
			want:     []int{2},
		},
		{
			commands: []string{"abc", "efg", "abcd"},
			pattern:  "abc",
			want:     []int{0, 2},
		},
		{
			commands: []string{"foo", "bar", "baz"},
			pattern:  "xyz",
			want:     []int{},
		},
		{
			commands: []string{"foo", "bar"},
			pattern:  "",
			want:     []int{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case_%d_%s", i, tt.pattern), func(t *testing.T) {
			matches := f.Match(tt.commands, tt.pattern)

			if len(matches) != len(tt.want) {
				t.Fatal(fmt.Sprintf("got %d matches | want %d matches", len(matches), len(tt.want)))
			}

			for i, match := range matches {
				if match.Index != tt.want[i] {
					t.Errorf("match[%d].Index = %d | want %d", i, match.Index, tt.want[i])
				}
			}
		})
	}
}
