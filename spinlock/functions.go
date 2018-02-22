package spinlock

// PartA finds the number immediately after the last insertion
func PartA(steps int) int {
	buffer := []int{0}
	currPos := 0
	for i := 1; i <= 2017; i++ {
		buffer, currPos = performSpin(buffer, currPos, steps, i)
	}
	return buffer[currPos+1]
}

func performSpin(buffer []int, currPos, steps, val int) ([]int, int) {
	currPos = (currPos + steps) % len(buffer)
	frontSlice := buffer[0 : currPos+1]
	backSlice := make([]int, len(buffer)-currPos-1)
	copy(backSlice, buffer[currPos+1:len(buffer)])
	frontSlice = append(frontSlice, val)
	buffer = append(frontSlice, backSlice...)
	currPos++
	return buffer, currPos
}
