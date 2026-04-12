package utils

import (
	"fmt"
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
