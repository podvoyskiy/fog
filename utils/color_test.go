package utils

import "testing"

func TestColorBuilder(t *testing.T) {
	tests := []struct {
		name  string
		color *ColorBuilder
		input string
		want  string
	}{
		{"red", Red(), "test", "\033[31mtest\033[0m"},
		{"blue bold", Blue().Bold(), "test", "\033[1;34mtest\033[0m"},
		{"cyan underline", Cyan().Underline(), "test", "\033[4;36mtest\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.color.Sprint(tt.input)
			if got != tt.want {
				t.Errorf("got : %q | want : %q", got, tt.want)
			}
		})
	}
}
