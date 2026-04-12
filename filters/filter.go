package filters

import (
	"fmt"
)

type FilterType uint8

const (
	typeSkim FilterType = iota + 1
	typeSubstring
)

func (f FilterType) uint8() uint8 {
	return uint8(f)
}

type MatchResult struct {
	Score int
	Index int
}

type Filter interface {
	GetId() uint8
	Match(commands []string, pattern string) []MatchResult
}

func Default() Filter {
	return &skimFilter{}
}

func FromUint8(id uint8) (Filter, error) {
	switch FilterType(id) {
	case typeSkim:
		return &skimFilter{}, nil
	case typeSubstring:
		return &substringFilter{}, nil
	default:
		return nil, fmt.Errorf("unknown filter type %d, expected %d (skim) or %d (substring)", id, typeSkim, typeSubstring)
	}
}
