package stringutils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if "olleh ,ih" != Reverse("hi, hello") {
		t.Errorf("reverse string was incorrect, expected : %v, found %v", "olleh ,ih", Reverse("hi, hello"))
	}
}

func TestEncode(t *testing.T) {
	if "b" != Encode("a", 1) {
		t.Errorf("encoded string was incorrect, expected : %v, found %v", "b", Encode("a", 1))
	}
}

func TestDecode(t *testing.T) {
	if "a" != Decode("b", 1) {
		t.Errorf("decoded string was incorrect, expected : %v, found %v", "a", Decode("b", 1))
	}
}
