package filters

import (
	"fmt"
)

type MatchResult struct {
	Score int
	Index int
}

type Filtering interface {
	GetId() uint8
	Match(commands []string, pattern string) []MatchResult
}

func Default() Filtering {
	return &filter{}
}

func FromUint8(id uint8) (Filtering, error) {
	switch FilterType(id) {
	case typeDefault:
		return &filter{}, nil
	case typeFuzzy:
		return &fuzzyFilter{}, nil
	case typeFrequency:
		return &FrequencyFilter{}, nil
	default:
		return nil, fmt.Errorf("unknown filter type %d, expected %s",
			id, availableFilters())
	}
}
