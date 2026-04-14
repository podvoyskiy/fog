package filters

// first tries exact substring matches sorted by frequency. if no exact matches found, falls back to fuzzy search.
type filter struct{}

var _ Filtering = (*filter)(nil)

func (f *filter) GetId() uint8 {
	return typeDefault.uint8()
}

func (f *filter) Match(commands []string, pattern string) []MatchResult {
	// priority 1: exact matches with frequency sorting
	if matches := (&FrequencyFilter{}).Match(commands, pattern); len(matches) > 0 {
		return matches
	}
	// priority 2: fuzzy search
	return (&fuzzyFilter{}).Match(commands, pattern)
}
