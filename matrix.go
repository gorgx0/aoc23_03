package main

import "log"

type numberWithPosition struct {
	number   string
	position position
}

func MakeMatrix(matrix []string) [][]string {
	result := make([][]string, len(matrix))
	for i, s := range matrix {
		result[i] = make([]string, len(s))
		for j, c := range s {
			result[i][j] = string(c)
		}
	}
	return result
}

func IsNumber(matrix [][]string, position position) bool {
	return isNumber(matrix[position.y][position.x])
}

func IsSymbol(matrix [][]string, position position) bool {
	return !isDotOrNumber(matrix[position.y][position.x])
}

func GetNumbersWithPositions(matrix [][]string) []numberWithPosition {
	var result []numberWithPosition
	var tmp numberWithPosition
	for y, row := range matrix {
		tmp = numberWithPosition{}
		x := 0
		for x < len(row) {
			if isNumber(matrix[y][x]) {
				if tmp.number == "" {
					tmp.position = position{x: x, y: y}
				}
				tmp.number += matrix[y][x]
			} else {
				if tmp.number != "" {
					result = append(result, tmp)
					tmp = numberWithPosition{}
				}
			}
			x++
		}
		if tmp.number != "" {
			result = append(result, tmp)
			continue
		}
		if isNumber(matrix[y][x-1]) {
			tmp.number = matrix[y][x-1]
			tmp.position = position{x: x - 1, y: y}
			result = append(result, tmp)
		}
	}
	return result
}

func FilterOutNumbersWithSymbols(matrix [][]string) []string {
	var result []string

	numbersWithPosition := GetNumbersWithPositions(matrix)
	log.Println(numbersWithPosition)

FOUND:
	for _, numberWithPosition := range numbersWithPosition {

		checkLeft := numberWithPosition.position.x-1 >= 0
		checkRight := numberWithPosition.position.x+len(numberWithPosition.number) < len(matrix[0])
		checkUp := numberWithPosition.position.y-1 >= 0
		checkDown := numberWithPosition.position.y+1 < len(matrix)

		yStart := numberWithPosition.position.y - 1
		xStart := numberWithPosition.position.x - 1
		yEnd := numberWithPosition.position.y + 1
		xEnd := numberWithPosition.position.x + len(numberWithPosition.number)

		if !checkLeft {
			xStart = 0
		}
		if !checkRight {
			xEnd = len(matrix[0]) - 1
		}
		if !checkUp {
			yStart = 0
		}
		if !checkDown {
			yEnd = len(matrix) - 1
		}

		if checkUp {
			for x := xStart; x <= xEnd; x++ {
				if IsSymbol(matrix, position{x: x, y: yStart}) {
					result = append(result, numberWithPosition.number)
					continue FOUND
				}
			}
		}

		if checkDown {
			for x := xStart; x <= xEnd; x++ {
				if IsSymbol(matrix, position{x: x, y: yEnd}) {
					result = append(result, numberWithPosition.number)
					continue FOUND
				}
			}
		}

		if checkLeft {
			for y := yStart; y <= yEnd; y++ {
				if IsSymbol(matrix, position{x: xStart, y: y}) {
					result = append(result, numberWithPosition.number)
					continue FOUND
				}
			}
		}

		if checkRight {
			for y := yStart; y <= yEnd; y++ {
				if IsSymbol(matrix, position{x: xEnd, y: y}) {
					result = append(result, numberWithPosition.number)
					continue FOUND
				}
			}
		}
	}
	return result
}
