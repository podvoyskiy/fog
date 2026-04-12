package search

import (
	f "github.com/podvoyskiy/fog/filters"
	"github.com/podvoyskiy/fog/history"
)

type Searcher struct {
	filter          f.Filter
	MaxResults      uint8
	Commands        []string
	FilteredIndices []int
	SelectedIndex   int
	SearchQuery     string
}

func Init(f f.Filter, maxResults uint8) (*Searcher, error) {
	history, err := history.Load()
	if err != nil {
		return nil, err
	}

	return &Searcher{
		filter:          f,
		MaxResults:      maxResults,
		Commands:        history.Commands,
		FilteredIndices: nil,
		SelectedIndex:   0,
		SearchQuery:     "",
	}, nil
}

func (s *Searcher) ApplyFilter() {
	s.SelectedIndex = 0

	matches := s.filter.Match(s.Commands, s.SearchQuery)
	limit := min(len(matches), int(s.MaxResults))

	indices := make([]int, 0, limit)
	for i := range limit {
		indices = append(indices, matches[i].Index)
	}

	s.FilteredIndices = indices
}

func (s *Searcher) GetSelectedCommand() (string, bool) {
	return s.GetCommandByIndex(s.SelectedIndex)
}

func (s *Searcher) GetCommandByIndex(index int) (string, bool) {
	if index < 0 || index >= len(s.FilteredIndices) {
		return "", false
	}

	idx := s.FilteredIndices[index]
	if idx < len(s.Commands) {
		return s.Commands[idx], true
	}

	return "", false
}

func (s *Searcher) ResultCount() int {
	return len(s.FilteredIndices)
}
