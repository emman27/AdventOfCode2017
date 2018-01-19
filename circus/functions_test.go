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

func TestFindUnbalanced(t *testing.T) {
	if res := (&Program{Weight: 1}).findUnbalancedNode(); res != nil {
		t.Fatalf("Failed test case")
	}
	if res := (&Program{Weight: 1, Children: []*Program{
		&Program{Weight: 10},
		&Program{Weight: 10},
		&Program{Weight: 1, Children: []*Program{
			&Program{Weight: 3},
			&Program{Weight: 3},
			&Program{Weight: 3},
		}},
	}}).findUnbalancedNode(); res != nil {
		t.Fatalf("Failed test case #2")
	}
	easyAnswer := &Program{Weight: 1}
	if res := (&Program{Weight: 1, Children: []*Program{
		easyAnswer,
		&Program{Weight: 2},
		&Program{Weight: 2},
	}}).findUnbalancedNode(); res != easyAnswer {
		t.Fatalf("Failed test case #3")
	}

	// Always find the deepest nested problem
	answer := &Program{Weight: 1}
	hardTestCase := &Program{Weight: 1, Children: []*Program{
		&Program{Weight: 10},
		&Program{Weight: 10},
		&Program{Weight: 4, Children: []*Program{
			answer,
			&Program{Weight: 2},
			&Program{Weight: 2},
		}},
	}}
	if res := hardTestCase.findUnbalancedNode(); res != answer {
		t.Fatalf("Failed test case #4")
	}

	evenHarderAnswer := &Program{Name: "ans", Weight: 1, Children: []*Program{
		&Program{Weight: 2},
		&Program{Weight: 2},
		&Program{Weight: 2},
	}}
	testCase3 := &Program{Weight: 1, Children: []*Program{
		&Program{Weight: 10, Name: "a"},
		evenHarderAnswer,
		&Program{Weight: 10, Name: "c"},
	}}
	if res := testCase3.findUnbalancedNode(); res != evenHarderAnswer {
		t.Fatalf("Failed test case #5, wanted %v but got %v", evenHarderAnswer, res)
	}

	twoBranch := &Program{Children: []*Program{
		&Program{Weight: 6, Children: []*Program{&Program{Weight: 6}}},
		&Program{Children: []*Program{
			&Program{Weight: 4},
			&Program{Weight: 5},
			&Program{Weight: 4},
		}},
	}}
	if res := twoBranch.findUnbalancedNode(); res.Weight != 5 {
		t.Fatalf("Two branch failed")
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
		level1.AddChild(p)
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
		l1.AddChild(p)
	}
	for _, p := range l3 {
		p.AddParent(unbalanced)
		unbalanced.AddChild(p)
	}
	if w, err := unbalanced.expectedWeight(); err != nil || w != 10 {
		t.Fatalf("Failed expectedWeight test 2, wanted %d, got %d", 10, w)
	}
	if unbalanced.childrenWeight() != 9 {
		t.Fatalf("Now I just feel stupid for getting this case wrong")
	}
	if l1.findUnbalancedNode() != unbalanced {
		t.Fatalf("Failed to find unbalanced node")
	}
}
