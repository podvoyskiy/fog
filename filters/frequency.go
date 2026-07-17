package filters

import (
	"strings"
)

type FrequencyFilter struct{}

var _ Filtering = &FrequencyFilter{}

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

	SortResults(results, commands)

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

	SortResults(results, commands)

	return results
}
