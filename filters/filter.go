package filters

import (
	"fmt"
)

type FilterType uint8

const (
	typeFuzzy FilterType = iota + 1
	typeSubstring
	typeFrequency
)

func GetFilterTypes() []FilterType {
	return []FilterType{typeFuzzy, typeSubstring, typeFrequency}
}

func (f FilterType) uint8() uint8 {
	return uint8(f)
}

func (f FilterType) toString() string {
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

type MatchResult struct {
	Score int
	Index int
}

type Filter interface {
	GetId() uint8
	GetName() string
	Match(commands []string, pattern string) []MatchResult
}

func Default() Filter {
	return &fuzzyFilter{}
}

func FromUint8(id uint8) (Filter, error) {
	switch FilterType(id) {
	case typeFuzzy:
		return &fuzzyFilter{}, nil
	case typeSubstring:
		return &substringFilter{}, nil
	case typeFrequency:
		return &frequencyFilter{}, nil
	default:
		return nil, fmt.Errorf("unknown filter type %d, expected %d (fuzzy) or %d (substring) or %d (frequency)",
			id, typeFuzzy, typeSubstring, typeFrequency)
	}
}
