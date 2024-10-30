package day12

import "testing"

func TestPart1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt")
	exp := 31
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver("input.txt")
	exp := 497
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
func TestPart2(t *testing.T) {
	act := Part2Solver("testdata/test1.txt")
	exp := 29
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver("input.txt")
	exp := 492
	if act != exp {
		t.Errorf("%d != %d", act, exp)
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
