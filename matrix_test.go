package main

import (
	"testing"
)

func TestFilterOutNumbersWithSymbols(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]string
		want   []string
	}{
		{
			"4x4 matrix without two digit numer at 1,1 and symbol above",
			[][]string{
				{".", "#", ".", "."},
				{".", "1", "2", "."},
				{".", ".", ".", "."},
				{".", ".", ".", "."},
			},
			[]string{"12"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FilterOutNumbersWithSymbols(test.matrix)
			if len(result) != len(test.want) {
				t.Errorf("FilterOutNumbersWithSymbols(%v) = %v, want %v", test.matrix, result, test.want)
			} else {
				for i, number := range result {
					if number != test.want[i] {
						t.Errorf("FilterOutNumbersWithSymbols(%v) = %v, want %v", test.matrix, result, test.want)
					}
				}
			}
		})
	}

}

func TestGetNumbersWithPositions(t *testing.T) {

	tests := []struct {
		name   string
		matrix [][]string
		want   []numberWithPosition
	}{
		{
			"3,3 matrix",
			[][]string{
				{".", ".", "."},
				{".", ".", "."},
				{".", ".", "."},
			},
			[]numberWithPosition{},
		}, {
			"3,3 matrix with number at 1,1",
			[][]string{
				{".", ".", "."},
				{".", "4", "."},
				{".", ".", "."},
			},
			[]numberWithPosition{{number: "4", position: position{x: 1, y: 1}}},
		}, {
			"3,3 matrix with numbers in each row",
			[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			[]numberWithPosition{
				{number: "123", position: position{x: 0, y: 0}},
				{number: "456", position: position{x: 0, y: 1}},
				{number: "789", position: position{x: 0, y: 2}},
			},
		}, {
			"3,3 matrix with number at 0,0 1,1 2,2",
			[][]string{
				{"1", ".", "."},
				{".", "2", "."},
				{".", ".", "3"},
			},
			[]numberWithPosition{
				{number: "1", position: position{x: 0, y: 0}},
				{number: "2", position: position{x: 1, y: 1}},
				{number: "3", position: position{x: 2, y: 2}},
			},
		}, {
			"3,3 matrix with number at 0,0 1,1 1,2",
			[][]string{
				{"1", "1", "."},
				{".", "2", "2"},
				{".", "3", "3"},
			},
			[]numberWithPosition{
				{number: "11", position: position{x: 0, y: 0}},
				{number: "22", position: position{x: 1, y: 1}},
				{number: "33", position: position{x: 1, y: 2}},
			},
		}, {
			"5x5 matrix",
			[][]string{
				{".", "2", "2", ".", "."},
				{".", ".", ".", ".", "."},
				{".", "4", "5", "#", "."},
				{".", ".", "7", "8", "9"},
				{"1", ".", "5", ".", "7"},
			},
			[]numberWithPosition{
				{number: "22", position: position{x: 1, y: 0}},
				{number: "45", position: position{x: 1, y: 2}},
				{number: "789", position: position{x: 2, y: 3}},
				{number: "1", position: position{x: 0, y: 4}},
				{number: "5", position: position{x: 2, y: 4}},
				{number: "7", position: position{x: 4, y: 4}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := GetNumbersWithPositions(test.matrix)
			if len(result) != len(test.want) {
				t.Errorf("GetNumbersWithSymbols(%v) = %v, want %v", test.matrix, result, test.want)
			} else {
				for i, numberPosition := range result {
					if numberPosition.number != test.want[i].number || numberPosition.position.x != test.want[i].position.x || numberPosition.position.y != test.want[i].position.y {
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
