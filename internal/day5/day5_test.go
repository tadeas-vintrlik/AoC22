package day5

import (
	_ "embed"
	"testing"
)

//go:embed testdata/test1.txt
var test1 string

func TestPart1(t *testing.T) {
	act := Part1Solver(test1)
	exp := "CMZ"
	if act != exp {
		t.Errorf("%s != %s", act, exp)
	}
}

func TestPart2(t *testing.T) {
	act := Part2Solver(test1)
	exp := "MCD"
	if act != exp {
		t.Errorf("%s != %s", act, exp)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Solver(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2Solver(input)
	}
}
