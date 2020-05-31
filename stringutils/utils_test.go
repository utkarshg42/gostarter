package stringutils

import (
	"testing"
)


func TestReverse(t *testing.T) {
	if "olleh ,ih" != Reverse("hi, hello") {
		t.Errorf("reverse string was incorrect, expected : %v, found %v", "olleh ,ih", Reverse("hi, hello"))
	}
}

