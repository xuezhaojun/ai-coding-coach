package encode_decode_strings

import (
	"reflect"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		{
			name: "basic strings",
			strs: []string{"hello", "world"},
		},
		{
			name: "empty list",
			strs: []string{},
		},
		{
			name: "single empty string",
			strs: []string{""},
		},
		{
			name: "multiple empty strings",
			strs: []string{"", "", ""},
		},
		{
			name: "strings with special characters",
			strs: []string{"he:llo", "wor#ld", "foo;bar"},
		},
		{
			name: "strings with delimiters and colons",
			strs: []string{"4:abcd", "3:xyz"},
		},
		{
			name: "single character strings",
			strs: []string{"a", "b", "c"},
		},
		{
			name: "mixed empty and non-empty",
			strs: []string{"", "a", "", "b", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := encode(tt.strs)
			decoded := decode(encoded)
			if decoded == nil && len(tt.strs) == 0 {
				return
			}
			if !reflect.DeepEqual(decoded, tt.strs) {
				t.Errorf("decode(encode(%v)) = %v, want %v", tt.strs, decoded, tt.strs)
			}
		})
	}
}
