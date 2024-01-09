package main

import (
	"testing"
)

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

func TestFindPart2GearNumbers(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]string
		want   []gear
	}{
		{
			name: "3x3 matrix with 0 gear",
			matrix: [][]string{
				{"3", ".", "5"},
				{".", "#", "."},
				{".", ".", "."},
			},
			want: []gear{},
		}, {
			name: "3x3 matrix with 1 gear",
			matrix: [][]string{
				{"4", ".", "."},
				{".", "*", "7"},
				{".", ".", "."},
			},
			want: []gear{
				{num0: "4", num1: "7"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPart2GearNumbers(tt.matrix); !equalGears(got, tt.want) {
				t.Errorf("FindPart2GearNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalGears(got []gear, want []gear) bool {
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
