package test

import (
	"testing"
)

func TestTest(t *testing.T) {
	if test() != "hello" {
		t.Errorf("WRONG!")
	}
	if test2() != "world" {
		t.Errorf("WRONG!")
	}
}