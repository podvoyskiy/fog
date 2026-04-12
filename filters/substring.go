package filters

import "strings"

type substringFilter struct{}

var _ Filter = (*substringFilter)(nil)

func (f *substringFilter) GetId() uint8 {
	return typeSubstring.uint8()
}

func (f *substringFilter) Match(commands []string, pattern string) []MatchResult {
	if pattern == "" {
		return nil
	}

	patternLower := strings.ToLower(pattern)
	var results []MatchResult

	for i, cmd := range commands {
		cmdLower := strings.ToLower(cmd)

		pos := strings.Index(cmdLower, patternLower) //find first occurrence position
		if pos != -1 {
			//score: 100 minus position (earlier occurrence = higher score)
			score := 100 - pos
			results = append(results, MatchResult{
				Score: score,
				Index: i,
			})
		}
	}

	return results
}
