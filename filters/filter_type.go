package filters

import (
	"fmt"
	"strings"
)

type FilterType uint8

const (
	typeHybrid FilterType = iota + 1
	typeFuzzy
	typeSubstring
	typeFrequency
)

func AllFilterTypes() []FilterType {
	return []FilterType{typeHybrid, typeFuzzy, typeSubstring, typeFrequency}
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
	case typeHybrid:
		return "hybrid(frequency then fuzzy)"
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
