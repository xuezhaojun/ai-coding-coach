package serialize_tree

import (
	"reflect"
	"testing"
)

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		{"example 1", []int{1, 2, 3, -101, -101, 4, 5}},
		{"empty tree", []int{}},
		{"single node", []int{1}},
		{"left only", []int{1, 2}},
		{"right only", []int{1, -101, 2}},
		{"full tree", []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			data := serialize(root)
			got := treeToSlice(deserialize(data))
			want := treeToSlice(root)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("serialize/deserialize roundtrip = %v, want %v", got, want)
			}
		})
	}
}
