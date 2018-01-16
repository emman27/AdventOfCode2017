package trampolines

// PartA returns the number of steps to get out of the maze
func PartA(steps []int) int {
	return jumper(steps, func(i int) int { return i + 1 })
}

// PartB returns the number of steps out of a maze with a custom jumper function
func PartB(steps []int) int {
	return jumper(steps, func(i int) int {
		if i >= 3 {
			return i - 1
		}
		return i + 1
	})
}

func jumper(steps []int, increment func(int) int) int {
	curr := 0
	i := 0
	for curr < len(steps) {
		tmp := curr
		curr = curr + steps[curr]
		steps[tmp] = increment(steps[tmp])
		i++
	}
	return i
}
