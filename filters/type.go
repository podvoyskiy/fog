package filters

import (
	"fmt"
	"strings"
)

type FilterType uint8

const (
	typeDefault FilterType = iota + 1
	typeFuzzy
	typeFrequency
)

func allFilterTypes() []FilterType {
	return []FilterType{typeDefault, typeFuzzy, typeFrequency}
}

func availableFilters() string {
	var parts []string
	for _, ft := range allFilterTypes() {
		parts = append(parts, fmt.Sprintf("%d - %s", ft.uint8(), ft))
	}
	return strings.Join(parts, ", ")
}

func (f FilterType) uint8() uint8 {
	return uint8(f)
}

func (f FilterType) String() string {
	switch f {
	case typeDefault:
		return "default"
	case typeFuzzy:
		return "fuzzy"
	case typeFrequency:
		return "frequency"
	default:
		return "unknown"
	}
}
