package day10

import (
	_ "embed"
	"testing"
)

//go:embed testdata/test1.txt
var test1 string

func TestPart1(t *testing.T) {
	act := Part1Solver(test1)
	exp := 13140
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

func TestPart1Final(t *testing.T) {
	act := Part1Solver(input)
	exp := 11960
	if act != exp {
		t.Errorf("%d != %d", act, exp)
	}
}

var test2 string = `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`

func TestPart2(t *testing.T) {
	act := "\n" + Part2Solver(test1)
	exp := test2
	if act != exp {
		t.Errorf("\n%v\n != \n%v\n", act, exp)
	}
}

var part2res = `
####...##..##..####.###...##..#....#..#.
#.......#.#..#.#....#..#.#..#.#....#..#.
###.....#.#....###..#..#.#....#....####.
#.......#.#....#....###..#.##.#....#..#.
#....#..#.#..#.#....#....#..#.#....#..#.
####..##...##..#....#.....###.####.#..#.`

func TestPart2Final(t *testing.T) {
	act := "\n" + Part2Solver(input)
	exp := part2res
	if act != exp {
		t.Errorf("\n%v\n != \n%v\n", act, exp)
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
