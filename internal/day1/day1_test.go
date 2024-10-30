package day1

import (
	_ "embed"
	"testing"
)

//go:embed testdata/test1.txt
var test1 string

func TestPart1(t *testing.T) {
	res := Part1Solver(test1)
	exp := 24000
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	res := Part2Solver(test1)
	exp := 45000
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}
