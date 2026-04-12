package filters

import (
	"sort"
	"strings"
)

type substringFilter struct{}

var _ Filter = (*substringFilter)(nil)

func (f *substringFilter) GetId() uint8 {
	return typeSubstring.uint8()
}

func (f *substringFilter) GetName() string {
	return typeSubstring.toString()
}

func (f *substringFilter) Match(commands []string, pattern string) []MatchResult {
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

		pos := strings.Index(cmd, pattern) //find first occurrence position
		if pos != -1 {
			//score: 100 minus position (earlier occurrence = higher score)
			score := 100 - pos
			results = append(results, MatchResult{
				Score: score,
				Index: i,
			})
		}
	}

	// sort by score (higher better), then by command length (shorter better)
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return len(commands[results[i].Index]) < len(commands[results[j].Index])
		}
		return results[i].Score > results[j].Score
	})

	return results
}
