package reconstruct_itinerary

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{
			name:    "standard itinerary",
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want:    []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name:    "lexical order matters",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name:    "single ticket",
			tickets: [][]string{{"JFK", "ATL"}},
			want:    []string{"JFK", "ATL"},
		},
		{
			name:    "round trip",
			tickets: [][]string{{"JFK", "ATL"}, {"ATL", "JFK"}},
			want:    []string{"JFK", "ATL", "JFK"},
		},
		{
			name:    "three cities chain",
			tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
			want:    []string{"JFK", "NRT", "JFK", "KUL"},
		},
		{
			name:    "duplicate tickets",
			tickets: [][]string{{"JFK", "ATL"}, {"ATL", "JFK"}, {"JFK", "ATL"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "ATL", "JFK", "ATL", "SFO"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findItinerary(tt.tickets)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary(%v) = %v, want %v", tt.tickets, got, tt.want)
			}
		})
	}
}
