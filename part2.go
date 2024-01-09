package main

import "log"

func FindStars(matrix [][]string) []position {
	var result []position
	for y, row := range matrix {
		for x, c := range row {
			if c == "*" {
				result = append(result, position{x: x, y: y})
			}
		}
	}
	return result
}

type gear struct {
	num0 string
	num1 string
}

func FindPart2GearNumbers(matrix [][]string) []gear {
	result := []gear{}

	numbersWithPositions := GetNumbersWithPositions(matrix)
	stars := FindStars(matrix)
	log.Println(stars)
	log.Println(numbersWithPositions)
	return result
}
