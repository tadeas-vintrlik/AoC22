package day18

import "testing"

func TestPart1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt")
	exp := 64
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver("input.txt")
	exp := 3564
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
