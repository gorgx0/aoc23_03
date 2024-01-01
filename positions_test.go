package main

import "testing"

func TestEmptyGetPositionsOfSymbols(t *testing.T) {
	matrix := [][]string{}
	positions := GetPositionsOfSymbols(matrix)
	if len(positions) != 0 {
		t.Errorf("Expected empty array, got %v", positions)
	}
}

func TestGetPositionsOfSymbols(t *testing.T) {
	matrix := [][]string{
		{".", "."},
		{".", "#"},
	}
	expected := []position{
		{x: 1, y: 1},
	}
	positions := GetPositionsOfSymbols(matrix)
	if len(positions) != 1 {
		t.Errorf("Expected 1 positions, got %v", positions)
	}
	if positions[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, positions)
	}
}

func TestGetPositionsOfSymbols2(t *testing.T) {
	testCases := []struct {
		matrix   [][]string
		expected []position
	}{
		{
			matrix: [][]string{
				{".", "."},
				{".", "#"},
			},
			expected: []position{
				{x: 1, y: 1},
			},
		},
	}

	for _, testCase := range testCases {
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
