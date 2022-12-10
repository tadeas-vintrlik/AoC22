package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	res := Part1Solver("testdata/test1.txt")
	exp := 24000
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestPart1Final(t *testing.T) {
	res := Part1Solver("input.txt")
	exp := 67016
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestPart2(t *testing.T) {
	res := Part2Solver("testdata/test1.txt")
	exp := 45000
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func TestPart2Final(t *testing.T) {
	res := Part2Solver("input.txt")
	exp := 200116
	if res != exp {
		t.Errorf("%d != %d", res, exp)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Solver("input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2Solver("input.txt")
	}
}
