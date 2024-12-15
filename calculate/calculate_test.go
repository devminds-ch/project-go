package calculate

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 4) != 5 {
		t.Fail()
	}
	if Sum(2, 4) != 5 {
		t.Fail()
	}
	if Sum(1, 4) != 5 {
		t.Fail()
	}
}

func TestSumNegative(t *testing.T) {
	if Sum(1, -1) != 0 {
		t.Fail()
	}
	if Sum(-1, 1) != 0 {
		t.Fail()
	}
	if Sum(-1, -1) != -2 {
		t.Fail()
	}
}
