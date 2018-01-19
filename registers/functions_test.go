package registers

import "testing"

func TestPartA(t *testing.T) {
	if res := PartA("./test_input.txt"); res != 1 {
		t.Fatalf("Failed. Wanted %d, Got %d", 1, res)
	}
}
