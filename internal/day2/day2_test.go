package day2

import (
	_ "embed"
	"testing"
)

//go:embed testdata/test1.txt
var test1 string

func TestPart1(t *testing.T) {
	act := Part1Solver(test1)
	exp := 15
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2(t *testing.T) {
	act := Part2Solver(test1)
	exp := 12
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
