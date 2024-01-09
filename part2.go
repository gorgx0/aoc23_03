package main

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
