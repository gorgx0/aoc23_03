package main

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

func isDotOrNumber(symbol string) bool {
	return symbol == "." || symbol == "0" || symbol == "1" || symbol == "2" || symbol == "3" || symbol == "4" || symbol == "5" || symbol == "6" || symbol == "7" || symbol == "8" || symbol == "9"
}
