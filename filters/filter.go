package filters

import (
	"fmt"
)

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
	return &hybridFilter{}
}

func FromUint8(id uint8) (Filter, error) {
	switch FilterType(id) {
	case typeHybrid:
		return &hybridFilter{}, nil
	case typeFuzzy:
		return &fuzzyFilter{}, nil
	case typeSubstring:
		return &substringFilter{}, nil
	case typeFrequency:
		return &frequencyFilter{}, nil
	default:
		return nil, fmt.Errorf("unknown filter type %d, expected %s",
			id, AvailableFilters())
	}
}
