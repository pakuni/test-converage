package main

import (
	"testing"
)

func TestMainUn(t *testing.T) {
	if aa() != "asd" {
		t.Log("pass1")
		t.Log("pass3");
	}
}

func TestMainEq(t *testing.T) {
	if aa() == "asd" {
		t.Log("pass2")
	}
}
