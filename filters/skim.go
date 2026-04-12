package filters

import (
	"github.com/sahilm/fuzzy"
)

type skimFilter struct{}

var _ Filter = (*skimFilter)(nil)

func (f *skimFilter) GetId() uint8 {
	return typeSkim.uint8()
}

func (f *skimFilter) Match(commands []string, pattern string) []MatchResult {
	if pattern == "" {
		return nil
	}

	matches := fuzzy.Find(pattern, commands)

	results := make([]MatchResult, len(matches))
	for i, m := range matches {
		//fuzzy.Find: lower score = better match
		//invert to make higher score = better match (0-100 range)
		score := min(m.Score, 100)
		normalized := 100 - score

		results[i] = MatchResult{
			Score: normalized,
			Index: m.Index,
		}
	}

	return results
}
