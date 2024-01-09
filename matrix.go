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

	for _, numberWithPosition := range numbersWithPosition {
		xLeft := numberWithPosition.position.x - 1
		xRight := numberWithPosition.position.x + len(numberWithPosition.number)
		yUp := numberWithPosition.position.y - 1
		//yDown := numberWithPosition.position.y + 1

		if yUp >= 0 {
			if xLeft >= 0 {
				for x := xLeft; x < xRight; x++ {
					if IsSymbol(matrix, position{x: x, y: yUp}) {
						result = append(result, numberWithPosition.number)
						break
					}
				}
			}
		}
	}
	return result
}
