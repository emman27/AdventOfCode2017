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
	if res := PartB("./data_test.txt"); res != 60 {
		t.Fatalf("Failed Part B, wanted: %d, got %d", 60, res)
	}
}

func TestExpectedWeight(t *testing.T) {
	level1 := &Program{}
	level2 := []*Program{
		&Program{Weight: 2},
		&Program{Weight: 1},
		&Program{Weight: 2},
	}
	for _, p := range level2 {
		p.AddParent(level1)
	}
	for _, p := range level2 {
		if w, err := p.expectedWeight(); w != 2 || err != nil {
			t.Fatalf("Expected weight case 1 failed")
		}
	}

	l1 := &Program{}
	unbalanced := &Program{}
	l2 := []*Program{
		&Program{Weight: 10},
		unbalanced,
		&Program{Weight: 10},
	}
	l3 := []*Program{
		&Program{Weight: 3},
		&Program{Weight: 3},
		&Program{Weight: 3},
	}
	for _, p := range l2 {
		p.AddParent(l1)
	}
	for _, p := range l3 {
		p.AddParent(unbalanced)
	}
	if w, err := unbalanced.expectedWeight(); err != nil || w != 10 {
		t.Fatalf("Failed expectedWeight test 2, wanted %d, got %d", 10, w)
	}
	if unbalanced.childrenWeight() != 9 {
		t.Fatalf("Now I just feel stupid for getting this case wrong")
	}
}
