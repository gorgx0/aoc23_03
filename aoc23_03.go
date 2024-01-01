package main

func main() {
	input, err := readInput("test.txt")
	if err != nil {
		panic(err)
	}

	matrix := make([][]string, len(input))

	for i, s := range input {
		matrix[i] = make([]string, len(s))
		for j, c := range s {
			matrix[i][j] = string(c)
		}
	}

	positions := getAllSymbolsPositions(matrix)
	println(positions)
}

type position struct {
	x int
	y int
}

func getAllSymbolsPositions(matrix [][]string) []position {
	var positions []position
	for i, row := range matrix {
		for j, symbol := range row {
			if symbol == "." {
				break
			}
			positions = append(positions, position{x: i, y: j})
		}
	}
	return positions
}
