package spiralMemory

// Storage stores the data
type Storage struct {
	Data map[int]map[int]int
}

// NewStorage creates a new storage object
func NewStorage() *Storage {
	return &Storage{make(map[int]map[int]int)}
}

// Set sets a value on a storage
func (s *Storage) Set(x, y, val int) {
	if s.Data[x] == nil {
		s.Data[x] = make(map[int]int)
	}
	s.Data[x][y] = val
}

// Get gets the value of a storage at a particular element
func (s *Storage) Get(x, y int) int {
	return s.Data[x][y]
}

// Calculate gives the sum of surrounding cells
func (s *Storage) Calculate(x, y int) int {
	return s.Get(x-1, y) + s.Get(x+1, y) + s.Get(x, y-1) + s.Get(x, y+1) + s.Get(x-1, y-1) + s.Get(x+1, y+1) + s.Get(x-1, y+1) + s.Get(x+1, y-1)
}
