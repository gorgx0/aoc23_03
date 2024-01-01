package main

import "strings"

type position struct {
	x int
	y int
}

func GetNumberStartingAtPosition(position position, matrix [][]string) string {
	var result string
	result = matrix[position.x][position.y]
	return result
}

func GetPositionsOfSymbols(matrix [][]string) []position {
	var positions []position
	for i, row := range matrix {
		for j, symbol := range row {
			if !isDotOrNumber(symbol) {
				positions = append(positions, position{x: i, y: j})
			}
		}
	}
	return positions
}

func GetPositionsWithNumbersNearSymbolPosition(symbolPosition position, matrix [][]string) []position {
	var result []position
	for i := symbolPosition.x - 1; i <= symbolPosition.x+1; i++ {
		for j := symbolPosition.y - 1; j <= symbolPosition.y+1; j++ {
			if i == symbolPosition.x && j == symbolPosition.y {
				continue
			}
			if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[i]) {
				continue
			}
			if isNumber(matrix[i][j]) {
				result = append(result, position{x: i, y: j})
			}
		}
	}
	return result
}

func GetNumberNearPosition(position position, matrix [][]string) string {
	var numbers []string
	for i := position.x - 1; i <= position.x+1; i++ {
		for j := position.y - 1; j <= position.y+1; j++ {
			if i == position.x && j == position.y {
				continue
			}
			if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[i]) {
				continue
			}
			if isNumber(matrix[i][j]) {
				numbers = append(numbers, matrix[i][j])
			}
		}
	}

	return strings.Join(numbers, "")
}

func isDotOrNumber(symbol string) bool {
	return symbol == "." || isNumber(symbol)
}

func isNumber(symbol string) bool {
	return symbol == "0" || symbol == "1" || symbol == "2" || symbol == "3" || symbol == "4" || symbol == "5" || symbol == "6" || symbol == "7" || symbol == "8" || symbol == "9"
}
