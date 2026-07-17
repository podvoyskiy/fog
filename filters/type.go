package filters

type FilterType uint8

const (
	typeFuzzy FilterType = iota + 1
	typeFrequency
)

func newFilter(typeF FilterType) Filtering {
	switch typeF {
	case typeFrequency:
		return &FrequencyFilter{}
	case typeFuzzy:
		return NewFuzzyFilter()
	default:
		return NewFuzzyFilter()
	}
}

func (f FilterType) uint8() uint8 {
	return uint8(f)
}

func (f FilterType) String() string {
	switch f {
	case typeFrequency:
		return "frequency"
	case typeFuzzy:
		return "fuzzy"
	default:
		return "unknown"
	}
}
