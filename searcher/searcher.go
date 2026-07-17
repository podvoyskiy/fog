package search

import (
	"sort"

	c "github.com/podvoyskiy/fog/config"
	f "github.com/podvoyskiy/fog/filters"
	"github.com/podvoyskiy/fog/history"
	u "github.com/podvoyskiy/fog/utils"
)

type Searcher struct {
	config          *c.AppConfig
	commands        []string
	freqMap         map[string]int
	filteredIndices []int
	SelectedIndex   int
	SearchQuery     string
}

func Init(config *c.AppConfig) (*Searcher, error) {
	history, err := history.Load()
	if err != nil {
		return nil, err
	}

	freqMap := make(map[string]int)
	for _, cmd := range history.Commands {
		freqMap[cmd]++
	}

	return &Searcher{
		config:          config,
		commands:        history.Commands,
		freqMap:         freqMap,
		filteredIndices: nil,
		SelectedIndex:   0,
		SearchQuery:     "",
	}, nil
}

func (s *Searcher) ApplyFilter() {
	s.SelectedIndex = 0

	matches := s.config.Filter.Match(s.commands, s.SearchQuery)

	s.applyFreqBonus(matches)

	// re-sort including bonus
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score > matches[j].Score
	})

	limit := min(len(matches), int(s.config.Limit))

	indices := make([]int, 0, limit)
	for i := range limit {
		indices = append(indices, matches[i].Index)
	}

	s.filteredIndices = indices
}

func (s *Searcher) GetSelectedCommand() (string, bool) {
	return s.GetCommandByIndex(s.SelectedIndex)
}

func (s *Searcher) GetCommandByIndex(index int) (string, bool) {
	if index < 0 || index >= len(s.filteredIndices) {
		return "", false
	}

	idx := s.filteredIndices[index]
	if idx < len(s.commands) {
		return s.commands[idx], true
	}

	return "", false
}

func (s *Searcher) ResultCount() int {
	return len(s.filteredIndices)
}

func (s *Searcher) applyFreqBonus(matches []f.MatchResult) {
	if s.config.MaxFreqBonus == 0 {
		return
	}
	for i := range matches {
		cmd := s.commands[matches[i].Index]
		freq := s.freqMap[cmd]
		bonus := u.CalcFreqBonus(freq, int(s.config.MaxFreqBonus))
		matches[i].Score += bonus
	}
}
