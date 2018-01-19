package registers

import "testing"

func TestPartA(t *testing.T) {
	if res := PartA("./test_input.txt"); res != 1 {
		t.Fatalf("Failed. Wanted %d, Got %d", 1, res)
	}
}

func TestPartB(t *testing.T) {
	if res := PartB("./test_input.txt"); res != 10 {
		t.Fatalf("Failed. Wanted %d, Got %d", 10, res)
	}
}
