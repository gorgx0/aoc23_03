package main

import (
	"testing"
)

func TestGetNumbersWithSymbols(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]string
		want   []string
	}{
		{
			"5x5 matrix",
			[][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", "4", "5", "#", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			[]string{"45"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetNumbersWithSymbols(test.matrix)
			if len(result) != len(test.want) {
				t.Errorf("GetNumbersWithSymbols(%v) = %v, want %v", test.matrix, result, test.want)
			} else {
				for i, number := range result {
					if number != test.want[i] {
						t.Errorf("GetNumbersWithSymbols(%v) = %v, want %v", test.matrix, result, test.want)
					}
				}
			}
		})
	}
}

func TestIsSymbol(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]string
		position position
		want     bool
	}{
		{
			"3x3 matrix with symbol at 2,2",
			[][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "#"},
			},
			position{x: 2, y: 2},
			true,
		}, {
			"3x3 matrix with number at 2,1",
			[][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", "2", "$"},
			},
			position{x: 2, y: 1},
			false,
		}, {
			"3x3 matrix with dot at 2,1",
			[][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "$"},
			},
			position{x: 2, y: 1},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsSymbol(test.matrix, test.position)
			if got != test.want {
				t.Errorf("IsSymbol(%v, %v) = %v, want %v", test.matrix, test.position, got, test.want)
			}
		})
	}
}

func TestMakeMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix []string
	}{
		{
			"5x5 matrix",
			[]string{
				"12345",
				"67890",
				"12345",
				"67890",
				"12345",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log("test: ", test.name)
			matrix := MakeMatrix(test.matrix)
			for _, line := range matrix {
				t.Log(line)
			}
		})
	}
}

func TestIsNumber(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]string
		position position
		want     bool
	}{
		{
			"1x1 matrix with number",
			[][]string{
				{"1"},
			},
			position{0, 0},
			true,
		}, {
			"1x1 matrix with dot",
			[][]string{
				{"."},
			},
			position{0, 0},
			false,
		}, {
			"3x3 matrix with number at 1,1",
			[][]string{
				{".", ".", "."},
				{".", "1", "."},
				{".", ".", "."},
			},
			position{1, 1},
			true,
		}, {
			"3x3 matrix with dot at 1,1",
			[][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
			position{1, 1},
			false,
		}, {
			"3x3 matrix with number at 1,0",
			[][]string{
				{".", "1", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
			position{1, 0},
			true,
		}, {
			"3x3 matrix with number at 0,1",
			[][]string{
				{".", ".", "."},
				{"1", ".", "."},
				{".", ".", "."},
			},
			position{0, 1},
			true,
		}, {
			"3x3 matrix with number at 2,1",
			[][]string{
				{".", ".", "."},
				{".", ".", "2"},
				{".", ".", "."},
			},
			position{2, 1},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := IsNumber(test.matrix, test.position)
			if got != test.want {
				t.Errorf("IsNumber(%v, %v) = %v, want %v", test.matrix, test.position, got, test.want)
			}
		})
	}
}
