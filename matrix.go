package main

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

func GetNumbersWithSymbols(matrix [][]string) []string {
	return []string{"45"}
}
