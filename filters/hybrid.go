package filters

// first tries exact substring matches sorted by frequency. if no exact matches found, falls back to fuzzy search.
type hybridFilter struct{}

var _ Filter = (*hybridFilter)(nil)

func (f *hybridFilter) GetId() uint8 {
	return typeHybrid.uint8()
}

func (f *hybridFilter) GetName() string {
	return typeHybrid.String()
}

func (f *hybridFilter) Match(commands []string, pattern string) []MatchResult {
	// priority 1: exact matches with frequency sorting
	if matches := (&frequencyFilter{}).Match(commands, pattern); len(matches) > 0 {
		return matches
	}
	// priority 2: fuzzy search
	return (&fuzzyFilter{}).Match(commands, pattern)
}
