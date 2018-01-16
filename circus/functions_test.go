package circus

import (
	"testing"
)

func TestPartA(t *testing.T) {
	if res := PartA("./data_test.txt"); res != "tknk" {
		t.Fatalf("Failed, Got %s, But wanted %s", res, "tknt")
	}
}
