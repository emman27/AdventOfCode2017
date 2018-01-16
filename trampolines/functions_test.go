package trampolines

import "testing"

func TestPartA(t *testing.T) {
	testCase := []int{0, 3, 0, 1, -3}
	if res := PartA(testCase); res != 5 {
		t.Fatalf("Wanted %d, Got %d", 5, res)
	}
}

func TestPartB(t *testing.T) {
	testCase := []int{0, 3, 0, 1, -3}
	if res := PartB(testCase); res != 10 {
		t.Fatalf("Wanted %d, Got %d", 10, res)
	}
}
