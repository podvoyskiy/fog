package filters

import (
	"sort"
	"strings"
)

type FrequencyFilter struct{}

var _ Filtering = (*FrequencyFilter)(nil)

func (f *FrequencyFilter) GetId() uint8 {
	return typeFrequency.uint8()
}

func (f *FrequencyFilter) Match(commands []string, pattern string) []MatchResult {
	if pattern == "" {
		return nil
	}
	pattern = strings.ToLower(pattern)

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

	f.sortResults(results, commands)

	return results
}

func (f *FrequencyFilter) All(commands []string) []MatchResult {
	freq := make(map[string]int)
	for _, cmd := range commands {
		cmd = strings.ToLower(cmd)
		freq[cmd]++
	}

	seen := make(map[string]bool)
	var results []MatchResult

	for i, cmd := range commands {
		cmd = strings.ToLower(cmd)
		if !seen[cmd] {
			seen[cmd] = true
			results = append(results, MatchResult{
				Score: freq[cmd],
				Index: i,
			})
		}
	}

	f.sortResults(results, commands)

	return results
}

// sort by frequency (higher = better). if score equal - by command length (shorter better)
func (f *FrequencyFilter) sortResults(results []MatchResult, commands []string) {
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return len(commands[results[i].Index]) < len(commands[results[j].Index])
		}
		return results[i].Score > results[j].Score
	})
}
