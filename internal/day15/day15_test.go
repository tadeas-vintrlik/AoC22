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
	act := Part1Solver("input.txt", 2000000)
	exp := 5240818
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2(t *testing.T) {
	act := Part2Solver("testdata/test1.txt", 20)
	exp := 56000011
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver("input.txt", 4000000)
	exp := 13213086906101
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1Solver("input.txt", 2000000)
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2Solver("input.txt", 4000000)
	}
}
