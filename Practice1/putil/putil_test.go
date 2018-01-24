package putil

import "testing"

func TestReverse(t *testing.T) {
	input := []uint8("abcdefg")
	expect := "gfedcba"
	reverse([]uint8(input))
	if expect != string(input) {
		t.Errorf("expect %s, got %s", expect, input)
	}
}

func TestStrConv(t *testing.T) {
	input := "10232"
	expect := int64(10232)
	result := StrConv([]byte(input))
	if expect != result {
		t.Errorf("expect %d, got %d", expect, result)
	}
}

func TestBaseConv(t *testing.T) {

}
