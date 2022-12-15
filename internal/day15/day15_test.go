package day15

import "testing"

func TestPart1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt", 10)
	exp := 26
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver("input.txt", 10)
	exp := 26
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
