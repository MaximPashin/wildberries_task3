package main

import (
	"testing"
)

func TestCorrect(t *testing.T) {
	res, _ := Unpack("a4bc2d5e")
	if res != "aaaabccddddde" {
		t.Errorf("Unpack(a4bc2d5e) = %v; want aaaabccddddde", res)
	}

	res, _ = Unpack("abcd")
	if res != "abcd" {
		t.Errorf("Unpack(abcd) = %v; want abcd", res)
	}

	res, _ = Unpack("")
	if res != "" {
		t.Errorf("Unpack() = \"%v\"; want \"\"", res)
	}

	res, _ = Unpack("qwe\\4\\5")
	if res != "qwe45" {
		t.Errorf("Unpack(qwe\\4\\5) = %v; want qwe45", res)
	}

	res, _ = Unpack("qwe\\45")
	if res != "qwe44444" {
		t.Errorf("Unpack(qwe\\45) = %v; want qwe44444", res)
	}

	res, _ = Unpack("qwe\\\\5")
	if res != "qwe\\\\\\\\\\" {
		t.Errorf("Unpack(qwe\\\\5) = %v; want qwe\\\\\\\\\\", res)
	}
}

func TestIncorrect(t *testing.T) {
	res, err := Unpack("45")
	if err == nil {
		t.Errorf("Unpack(45) = %v; want error", res)
	}

	res, err = Unpack("abc\\")
	if err == nil {
		t.Errorf("Unpack(abc\\) = %v; want error", res)
	}
}
