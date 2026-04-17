package surrounded_regions

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{
			name: "standard case",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name:  "single cell X",
			board: [][]byte{{'X'}},
			want:  [][]byte{{'X'}},
		},
		{
			name:  "single cell O",
			board: [][]byte{{'O'}},
			want:  [][]byte{{'O'}},
		},
		{
			name: "all O on border",
			board: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O'},
				{'O', 'O'},
			},
		},
		{
			name: "O connected to border",
			board: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'O', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "surrounded O in center",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}
