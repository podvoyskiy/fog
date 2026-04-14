package filters

import (
	"sort"
	"strings"
)

type frequencyFilter struct{}

var _ Filtering = (*frequencyFilter)(nil)

func (f *frequencyFilter) GetId() uint8 {
	return typeFrequency.uint8()
}

func (f *frequencyFilter) Match(commands []string, pattern string) []MatchResult {
	if pattern == "" {
		return nil
	}
	pattern = strings.ToLower(pattern)

	//calc frequency
	freq := make(map[string]int)
	for _, cmd := range commands {
		cmd = strings.ToLower(cmd)
		if strings.Contains(cmd, pattern) {
			freq[cmd]++
		}
	}

	seen := make(map[string]bool)
	var results []MatchResult

	for i, cmd := range commands {
		cmd = strings.ToLower(cmd)
		if strings.Contains(cmd, pattern) && !seen[cmd] {
			seen[cmd] = true
			results = append(results, MatchResult{
				Score: freq[cmd],
				Index: i,
			})
		}
	}

	//sort by frequency (higher = better). if score equal - by command length (shorter better)
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return len(commands[results[i].Index]) < len(commands[results[j].Index])
		}
		return results[i].Score > results[j].Score
	})

	return results
}
