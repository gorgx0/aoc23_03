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

	sum := GetSumOfNumbersNearSymbols(matrix)
	println(sum)
}
