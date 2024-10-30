package day16

import "testing"

func TestPart1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt")
	exp := 1651
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
