package partition_labels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []int
	}{
		{"example 1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"single char", "a", []int{1}},
		{"all same", "aaaa", []int{4}},
		{"all unique", "abcdef", []int{1, 1, 1, 1, 1, 1}},
		{"two partitions", "aabbb", []int{2, 3}},
		{"example 2", "eccbbbbdec", []int{10}},
		{"interleaved", "abac", []int{3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partitionLabels(tt.s)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("partitionLabels(%q) = %v, want %v", tt.s, got, tt.expected)
			}
		})
	}
}
