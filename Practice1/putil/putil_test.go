package putil

import (
	"reflect"
	"testing"
)

// ======== Public ========

func TestStrConv(t *testing.T) {
	input := "10232"
	expect := int64(10232)
	result := StrConv([]byte(input))
	if expect != result {
		t.Errorf("expect %d, got %d", expect, result)
	}
}

func TestBaseConv(t *testing.T) {
	input := []reflect.Value{
		reflect.ValueOf(int64(2)),
		reflect.ValueOf(int64(2047)),
	}
	expect := []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	result := reflect.ValueOf(BaseConv).Call(input)[0].Interface().([]uint8)
	if !reflect.DeepEqual(expect, result) {
		t.Errorf("expect %v, got %v", expect, result)
	}
}

func TestEncodeToks(t *testing.T) {
	input := []uint8{9, 13, 15, 0, 14, 10, 1, 15, 11, 5, 3, 2, 10, 4}
	expect := "9df0ea1fb532a4"
	result := EncodeToks(input)
	if expect != result {
		t.Errorf("expect %s, got %s", expect, result)
	}
}

// ======== Private ========

func TestReverse(t *testing.T) {
	input := []uint8("abcdefg")
	expect := "gfedcba"
	reverse([]uint8(input))
	if expect != string(input) {
		t.Errorf("expect %s, got %s", expect, input)
	}
}

func TestSerial(t *testing.T) {
	input := []reflect.Value{
		reflect.ValueOf(byte('a')),
		reflect.ValueOf(uint8(11)),
	}
	expect := "abcdefghijk"
	result := reflect.ValueOf(serial).Call(input)[0].Interface().(string)
	if expect != result {
		t.Errorf("expect %s, got %s", expect, result)
	}
}
