package day8

import (
	_ "embed"
	"testing"
)

//go:embed testadata/test1.txt
var test1 string

func TestPart1(t *testing.T) {
	act := Part1Solver(test1)
	exp := 21
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver(input)
	exp := 1796
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2(t *testing.T) {
	act := Part2Solver(test1)
	exp := 8
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver(input)
	exp := 288120
	if act != exp {
		t.Errorf("%d != %d", act, exp)
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
