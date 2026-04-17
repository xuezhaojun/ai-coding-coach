package daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{"typical case", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"decreasing", []int{30, 20, 10}, []int{0, 0, 0}},
		{"increasing", []int{10, 20, 30}, []int{1, 1, 0}},
		{"single element", []int{50}, []int{0}},
		{"all same", []int{70, 70, 70, 70}, []int{0, 0, 0, 0}},
		{"two elements warmer", []int{30, 60}, []int{1, 0}},
		{"two elements cooler", []int{60, 30}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
