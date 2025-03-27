package test

import (
	"InterviewPrep/src"
	"testing"
)

func TestHello(t *testing.T) {
	if src.Hello() != "Hello World" {
		t.Error("hello() != Hello World")
	}
}

func TestReverse(t *testing.T) {
	if src.Reverse("Hello World") != "dlroW olleH" {
		t.Error("Reverse() != dlroW olleH")
	}
}
