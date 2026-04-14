package search

import (
	f "github.com/podvoyskiy/fog/filters"
	"github.com/podvoyskiy/fog/history"
)

type Searcher struct {
	filter          f.Filtering
	limit           uint8
	commands        []string
	filteredIndices []int
	SelectedIndex   int
	SearchQuery     string
}

func Init(f f.Filtering, limit uint8) (*Searcher, error) {
	history, err := history.Load()
	if err != nil {
		return nil, err
	}

	return &Searcher{
		filter:          f,
		limit:           limit,
		commands:        history.Commands,
		filteredIndices: nil,
		SelectedIndex:   0,
		SearchQuery:     "",
	}, nil
}

func (s *Searcher) ApplyFilter() {
	s.SelectedIndex = 0

	matches := s.filter.Match(s.commands, s.SearchQuery)
	limit := min(len(matches), int(s.limit))

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
