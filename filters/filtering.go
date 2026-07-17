package filters

import "sort"

type MatchResult struct {
	Score int
	Index int
}

type Filtering interface {
	GetId() uint8
	Match(commands []string, pattern string) []MatchResult
}

func Default() Filtering {
	return NewFuzzyFilter()
}

// sort (higher = better). if score equal - by command length (shorter better)
func SortResults(results []MatchResult, commands []string) {
	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return len(commands[results[i].Index]) < len(commands[results[j].Index])
		}
		return results[i].Score > results[j].Score
	})
}
