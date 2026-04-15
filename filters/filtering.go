package filters

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
