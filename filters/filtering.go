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
	return &Filter{}
}

func FromUint8(id uint8) (Filtering, error) {
	switch FilterType(id) {
	case typeDefault:
		return &Filter{}, nil
	case typeFuzzy:
		return &fuzzyFilter{}, nil
	case typeFrequency:
		return &frequencyFilter{}, nil
	default:
		return nil, fmt.Errorf("unknown filter type %d, expected %s",
			id, availableFilters())
	}
}
