package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Uint8(s string) (uint8, error) {
	s = strings.TrimSpace(s)
	val, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", s, err)
	}
	return uint8(val), nil
}

func Must[T any](val T, err error) T {
	if err != nil {
		Red().Bold().Println(err)
		os.Exit(1)
	}
	return val
}

// CalcFreqBonus returns frequency bonus capped by maxFreqBonus.
//
//	freq:  1   2   5   10   26   50   101
//	bonus: 0   1   2   3    5    7    10   (maxFreqBonus=10)
func CalcFreqBonus(freq int, maxFreqBonus int) int {
	if freq <= 1 {
		return 0
	}
	bonus := int(math.Sqrt(float64(freq - 1)))
	if bonus > maxFreqBonus {
		bonus = maxFreqBonus
	}
	return bonus
}
