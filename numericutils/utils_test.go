package numericutils

import "testing"

func TestAbs(t *testing.T) {
	if 1 != Abs(-1) {
		t.Errorf("incorrect absolute value of %v, expected %v, found %v", -1, 1, Abs(-1))
	}
}

func TestSum(t *testing.T) {
	if 3 != Sum(2, 1) {
		t.Errorf("incorrect sum value of %v and %v, expected %v, found %v", 1, 2, 3, Sum(1, 2))
	}
}
