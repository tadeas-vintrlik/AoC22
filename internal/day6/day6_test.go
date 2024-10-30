package day6

import "testing"

var test1 string = "bvwbjplbgvbhsrlpgdmjqwftvncz"
var test2 string = "nppdvjthqldpwncqszvftbrmjlhg"
var test3 string = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
var test4 string = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

func Test1Part1(t *testing.T) {
	act := Part1Solver(test1)
	exp := 5
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test2Part1(t *testing.T) {
	act := Part1Solver(test2)
	exp := 6
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test3Part1(t *testing.T) {
	act := Part1Solver(test3)
	exp := 10
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test4Part1(t *testing.T) {
	act := Part1Solver(test4)
	exp := 11
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver(input)
	exp := 1134
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

var test5 string = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
var test6 string = "bvwbjplbgvbhsrlpgdmjqwftvncz"
var test7 string = "nppdvjthqldpwncqszvftbrmjlhg"
var test8 string = "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"
var test9 string = "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

func Test1Part2(t *testing.T) {
	act := Part2Solver(test5)
	exp := 19
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test2Part2(t *testing.T) {
	act := Part2Solver(test6)
	exp := 23
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test3Part2(t *testing.T) {
	act := Part2Solver(test7)
	exp := 23
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test4Part2(t *testing.T) {
	act := Part2Solver(test8)
	exp := 29
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func Test5Part2(t *testing.T) {
	act := Part2Solver(test9)
	exp := 26
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart2Final(t *testing.T) {
	act := Part2Solver(input)
	exp := 2263
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
