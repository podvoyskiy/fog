package filters

import (
	"fmt"
	"strings"
)

type FilterType uint8

const (
	typeFuzzy FilterType = iota + 1
	typeSubstring
	typeFrequency
)

func AllFilterTypes() []FilterType {
	return []FilterType{typeFuzzy, typeSubstring, typeFrequency}
}

func AvailableFilters() string {
	var parts []string
	for _, ft := range AllFilterTypes() {
		parts = append(parts, fmt.Sprintf("%d - %s", ft.uint8(), ft))
	}
	return strings.Join(parts, ", ")
}

func (f FilterType) uint8() uint8 {
	return uint8(f)
}

func (f FilterType) String() string {
	switch f {
	case typeFuzzy:
		return "fuzzy"
	case typeSubstring:
		return "substring"
	case typeFrequency:
		return "frequency"
	default:
		return "unknown"
	}
}
