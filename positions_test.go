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
		{".", "#"},
		{".", "."},
	}
	positions := GetPositionsOfSymbols(matrix)
	if len(positions) != 1 {
		t.Errorf("Expected 1 positions, got %v", positions)
	}
}
