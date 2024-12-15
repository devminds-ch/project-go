package calculate

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 1) != 2 {
		t.Fail()
	}
	if Sum(2, 1) != 3 {
		t.Fail()
	}
	if Sum(1, 2) != 3 {
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

func TestSubtract(t *testing.T) {
	if Subtract(1, 1) != 0 {
		t.Fail()
	}
	if Subtract(2, 1) != 1 {
		t.Fail()
	}
	if Subtract(1, 2) != -1 {
		t.Fail()
	}
}

func TestSubtractNegative(t *testing.T) {
	if Subtract(1, -1) != 2 {
		t.Fail()
	}
	if Subtract(-1, 1) != -2 {
		t.Fail()
	}
	if Subtract(-1, -1) != 0 {
		t.Fail()
	}
}
