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
		&Program{Weight: 2},
		&Program{Weight: 2},
		&Program{Weight: 2},
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
		&Program{Weight: 4, Children: []*Program{
			answer,
			&Program{Weight: 2},
			&Program{Weight: 2},
		}},
		&Program{Weight: 10},
		&Program{Weight: 10},
	}}
	if res := hardTestCase.findUnbalancedNode(); res != answer {
		t.Fatalf("Failed test case #4")
	}
}
