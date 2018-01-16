package memoryReallocation

import (
	set "gopkg.in/fatih/set.v0"
)

// PartA calculates the number of cycles before detecting an infinite loop
func PartA(blocks []int) int {
	s := set.New()
	for i := 0; true; i++ {
		if s.Has(str(blocks)) {
			return i
		}
		s.Add(str(blocks))
		blocks = Reallocate(blocks)
	}
	return -1
}

// Reallocate moves the largest stack into the smaller stacks
func Reallocate(blocks []int) []int {
	idx, maximum := max(blocks)
	blocks[idx] = 0
	for i := idx + 1; maximum > 0; i++ {
		if i >= len(blocks) {
			i = i - len(blocks)
		}
		blocks[i]++
		maximum--
	}
	return blocks
}

func max(blocks []int) (int, int) {
	idx := 0
	max := blocks[0]
	for i, val := range blocks {
		if val > max {
			max = val
			idx = i
		}
	}
	return idx, max
}

func str(blocks []int) string {
	s := ""
	for _, num := range blocks {
		s += string(num) + ", "
	}
	return s
}
