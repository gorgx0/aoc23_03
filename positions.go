package main

func GetPositionsOfSymbols(matrix [][]string) []position {
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
