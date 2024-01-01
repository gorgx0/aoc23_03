package main

import "testing"

func TestEmptyGetPositionsOfSymbols(t *testing.T) {
	matrix := [][]string{}
	positions := GetPositionsOfSymbols(matrix)
	if len(positions) != 0 {
		t.Errorf("Expected empty array, got %v", positions)
	}
}

func TestGetPositionsOfSymbols2(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   [][]string
		expected []position
	}{
		{
			name:     "empty matrix",
			matrix:   nil,
			expected: nil,
		},
		{
			name: "single no symbol",
			matrix: [][]string{
				{"."},
			},
			expected: nil,
		},
		{
			name: "single number",
			matrix: [][]string{
				{"7"},
			},
			expected: nil,
		},
		{
			name: "single symbol",
			matrix: [][]string{
				{"#"},
			},
			expected: []position{
				{x: 0, y: 0},
			},
		},
		{
			name: "one symbol",
			matrix: [][]string{
				{".", "."},
				{".", "#"},
			},
			expected: []position{
				{x: 1, y: 1},
			},
		},
		{
			name: "bigger matrix",
			matrix: [][]string{
				{".", ".", "."},
				{".", "#", "."},
				{".", ".", "."},
			},
			expected: []position{
				{x: 1, y: 1},
			},
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)
		positions := GetPositionsOfSymbols(testCase.matrix)
		if len(positions) != len(testCase.expected) {
			t.Errorf("Expected %v positions, got %v", len(testCase.expected), len(positions))
		}
		for i, position := range positions {
			if position != testCase.expected[i] {
				t.Errorf("Expected %v, got %v", testCase.expected, positions)
			}
		}
	}
}
