package corruptionChecksum

// Calculate generates the checksum by finding the maximum difference between the smallest and largest values for each row, then summing these differences
func Calculate(matrix [][]int, calculatorFunction func([]int) int) int {
	var total int
	for _, row := range matrix {
		total += calculatorFunction(row)
	}
	return total
}

// MaxDiff calculates the maximum diff in a row
func MaxDiff(row []int) int {
	max := -2 ^ 21
	min := 2 ^ 21
	for _, num := range row {
		if max < num {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max - min
}

// EvenDivisResults finds the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.
func EvenDivisResults(row []int) int {
	return 0
}
