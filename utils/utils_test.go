package utils

import (
	"fmt"
	"testing"
)

func TestUint8(t *testing.T) {
	tests := []struct {
		input   string
		want    uint8
		wantErr bool
	}{
		{"123", 123, false},
		{" 42 ", 42, false},
		{"0", 0, false},
		{"255", 255, false},
		{"256", 0, true},
		{"-5", 0, true},
		{"abc", 0, true},
		{"", 0, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case_%s", tt.input), func(t *testing.T) {
			got, err := Uint8(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Uint8(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf("Uint8(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCalcFreqBonus(t *testing.T) {
	tests := []struct {
		freq         int
		maxFreqBonus int
		want         int
	}{
		{10, 0, 0},
		{0, 10, 0},
		{1, 10, 0},
		{2, 10, 1},
		{5, 10, 2},
		{10, 10, 3},
		{26, 10, 5},
		{50, 10, 7},
		{101, 10, 10},
		{300, 10, 10},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("case_%d_%d_%d", tt.freq, tt.maxFreqBonus, tt.want), func(t *testing.T) {
			got := CalcFreqBonus(tt.freq, tt.maxFreqBonus)
			if got != tt.want {
				t.Errorf("CalcFreqBonus(%d, %d) = %d, want %d", tt.freq, tt.maxFreqBonus, got, tt.want)
			}
		})
	}
}
