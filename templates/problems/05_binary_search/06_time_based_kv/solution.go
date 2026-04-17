package time_based_kv

import "sort"

// SolveTimeMap implements a time-based key-value store. Set is O(1), Get is O(log n). O(n) space.
type SolveTimeMap struct {
	store map[string][]entry
}

type entry struct {
	timestamp int
	value     string
}

func NewSolveTimeMap() SolveTimeMap {
	return SolveTimeMap{store: make(map[string][]entry)}
}

func (t *SolveTimeMap) Set(key string, value string, timestamp int) {
	t.store[key] = append(t.store[key], entry{timestamp, value})
}

func (t *SolveTimeMap) Get(key string, timestamp int) string {
	entries := t.store[key]
	if len(entries) == 0 {
		return ""
	}
	// Binary search for the largest timestamp <= given timestamp
	idx := sort.Search(len(entries), func(i int) bool {
		return entries[i].timestamp > timestamp
	})
	if idx == 0 {
		return ""
	}
	return entries[idx-1].value
}
