package circus

import (
	"testing"
)

func TestPartA(t *testing.T) {
	if res := PartA("./data_test.txt"); res != "tknk" {
		t.Fatalf("Failed, Got %s, But wanted %s", res, "tknt")
	}
}

func TestTotalWeight(t *testing.T) {
	child := Program{Weight: 1000}
	parent := Program{Weight: 50, Children: []*Program{&child}}
	if res := child.TotalWeight(); res != 1000 {
		t.Fatalf("Child weight wrong, got %d, want %d", res, 1000)
	}
	if res := parent.TotalWeight(); res != 1050 {
		t.Fatalf("Parent weight wrong, got %d, want %d", res, 1050)
	}
}

func TestPartB(t *testing.T) {
	if res := PartB("./data_test.txt"); res != 243 {
		t.Fatalf("Failed Part B, wanted: %d, got %d", 243, res)
	}
}
