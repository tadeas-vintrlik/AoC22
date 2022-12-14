package day14

import "testing"

func TestPart1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt")
	exp := 24
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver("input.txt")
	exp := 719
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}
func TestPart2(t *testing.T) {
	act := Part2Solver("testdata/test1.txt")
	exp := 93
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver("input.txt")
	exp := 23390
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
