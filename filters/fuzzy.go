package filters

import (
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type fuzzyFilter struct{}

var _ Filtering = (*fuzzyFilter)(nil)

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

		if fuzzy.Match(pattern, cmd) {
			rank := fuzzy.RankMatch(pattern, cmd)
			results = append(results, MatchResult{
				Score: rank,
				Index: i,
			})
		}
	}

	//fuzzy.Match returns lower score for better matches, so sort ascending
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score < results[j].Score
	})

	return results
}
