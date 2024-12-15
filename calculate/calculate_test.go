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

func TestMultiply(t *testing.T) {
	if Multiply(1, 1) != 1 {
		t.Fail()
	}
	if Multiply(2, 1) != 2 {
		t.Fail()
	}
	if Multiply(1, 2) != 2 {
		t.Fail()
	}
}

func TestMultiplyNegative(t *testing.T) {
	if Multiply(1, -1) != -1 {
		t.Fail()
	}
	if Multiply(-1, 1) != -1 {
		t.Fail()
	}
	if Multiply(-1, -1) != 1 {
		t.Fail()
	}
}
