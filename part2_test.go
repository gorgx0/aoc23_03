package main

import "testing"

func TestFindStars(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]string
		want   []position
	}{
		{
			name: "3x3 matrix with 0 star",
			matrix: [][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
			want: []position{},
		}, {
			name: "3x3 matrix with 1 star",
			matrix: [][]string{
				{".", ".", "."},
				{".", "*", "."},
				{".", ".", "."},
			},
			want: []position{
				{x: 1, y: 1},
			},
		}, {
			name: "3x3 matrix with 5 stars",
			matrix: [][]string{
				{"*", ".", "*"},
				{"*", "*", "."},
				{".", ".", "*"},
			},
			want: []position{
				{x: 0, y: 0},
				{x: 2, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
				{x: 2, y: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindStars(tt.matrix); !equalPositions(got, tt.want) {
				t.Errorf("FindStars() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalPositions(got []position, want []position) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
