package filters

import (
	"strings"

	fz "github.com/podvoyskiy/fuzzymatch"
)

type fuzzyFilter struct {
	matcher fz.FuzzyMatcher
}

var _ Filtering = &fuzzyFilter{}

func NewFuzzyFilter() *fuzzyFilter {
	return &fuzzyFilter{matcher: fz.NewMatcher()}
}

func (f *fuzzyFilter) GetId() uint8 {
	return typeFuzzy.uint8()
}

func (f *fuzzyFilter) Match(commands []string, pattern string) []MatchResult {
	if pattern == "" {
		return nil
	}
	pattern = strings.ToLower(pattern)

	seen := make(map[string]bool)
	var results []MatchResult

	for i, cmd := range commands {
		cmd = strings.ToLower(cmd)
		if seen[cmd] {
			continue
		}
		seen[cmd] = true

		if score, ok := f.matcher.FuzzyMatch(cmd, pattern); ok {
			results = append(results, MatchResult{
				Score: score,
				Index: i,
			})
		}
	}

	SortResults(results, commands)

	return results
}
