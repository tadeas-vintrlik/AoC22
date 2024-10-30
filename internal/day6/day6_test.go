package day6

import "testing"

func Test1Part1(t *testing.T) {
	act := Part1Solver("testdata/test1.txt")
	exp := 5
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test2Part1(t *testing.T) {
	act := Part1Solver("testdata/test2.txt")
	exp := 6
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test3Part1(t *testing.T) {
	act := Part1Solver("testdata/test3.txt")
	exp := 10
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test4Part1(t *testing.T) {
	act := Part1Solver("testdata/test4.txt")
	exp := 11
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver("input.txt")
	exp := 1134
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test1Part2(t *testing.T) {
	act := Part2Solver("testdata/test5.txt")
	exp := 19
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test2Part2(t *testing.T) {
	act := Part2Solver("testdata/test6.txt")
	exp := 23
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test3Part2(t *testing.T) {
	act := Part2Solver("testdata/test7.txt")
	exp := 23
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test4Part2(t *testing.T) {
	act := Part2Solver("testdata/test8.txt")
	exp := 29
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test5Part2(t *testing.T) {
	act := Part2Solver("testdata/test9.txt")
	exp := 26
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver("input.txt")
	exp := 2263
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
