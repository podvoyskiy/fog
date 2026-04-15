package filters

type FilterType uint8

const (
	typeDefault FilterType = iota + 1
	typeFuzzy
	typeFrequency
)

func newFilter(typeF FilterType) Filtering {
	switch typeF {
	case typeFuzzy:
		return &fuzzyFilter{}
	case typeFrequency:
		return &FrequencyFilter{}
	default:
		return &filter{}
	}
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
