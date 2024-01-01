package main

import (
	"testing"
)

func TestGetPositionsWithNumbersNearSymbolPosition(t *testing.T) {
	tc := []struct {
		name     string
		matrix   [][]string
		position position
		expected []position
	}{
		{
			name:     "empty matrix",
			matrix:   [][]string{},
			position: position{x: 0, y: 0},
			expected: nil,
		}, {
			name: "single number",
			matrix: [][]string{
				{".", "2"},
				{".", "."},
			},
			position: position{x: 0, y: 0},
			expected: []position{
				{x: 0, y: 1},
			},
		}, {
			name: "two numbers",
			matrix: [][]string{
				{".", "2"},
				{"3", "."},
			},
			position: position{x: 0, y: 0},
			expected: []position{
				{x: 0, y: 1},
				{x: 1, y: 0},
			},
		},
		{
			name: "more numbers in 3x3 matrix",
			matrix: [][]string{
				{"1", "2", "3"},
				{"4", "#", "6"},
				{"7", "8", "9"},
			},
			position: position{x: 1, y: 1},
			expected: []position{
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 0, y: 2},
				{x: 1, y: 0},
				{x: 1, y: 2},
				{x: 2, y: 0},
				{x: 2, y: 1},
				{x: 2, y: 2},
			},
		},
	}

	for _, testCase := range tc {
		t.Log(testCase.name)
		positions := GetPositionsWithNumbersNearSymbolPosition(testCase.position, testCase.matrix)
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

func TestGetNumberNearPosition(t *testing.T) {
	testCases := []struct {
		name           string
		matrix         [][]string
		position       position
		expectedNumber string
	}{
		{
			name: "empty matrix",
			matrix: [][]string{
				{},
			},
			position: position{x: 0, y: 0},
		},
		{
			name: "2x2 matrix with one number at position 0,1",
			matrix: [][]string{
				{"1", "."},
				{".", "."},
			},
			position:       position{x: 0, y: 1},
			expectedNumber: "1",
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.name)
		number := GetNumberNearPosition(testCase.position, testCase.matrix)
		if number != testCase.expectedNumber {
			t.Errorf("Expected %v, got %v", testCase.expectedNumber, number)
		}
	}
}

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
				{"@", ".", "."},
				{".", "#", "."},
				{".", ".", "$"},
			},
			expected: []position{
				{x: 0, y: 0},
				{x: 1, y: 1},
				{x: 2, y: 2},
			},
		}, {
			name: "empty 5x5 matrix",
			matrix: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			},
			expected: []position{},
		}, {
			name: "5x5 matrix with symbols",
			matrix: [][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", "#", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
			}, expected: []position{
				{x: 2, y: 2},
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
