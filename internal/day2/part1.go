package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type round struct {
	you  byte
	them byte
}

// X => A  B  C
// Y => B  C  A
// Z => C  A  B
// The order is tie, lose, win this was chosen for X and A to be first for a prettier mapping function
// You might notice that each line is same except for offset by one
var round_result = []int{3, 0, 6}

func (r round) getScore() int {
	// Get the value for your play (distance from X) and score for round result
	return int(r.you-'X'+1) + round_result[(2*int((r.you-'X'))+int((r.them-'A')))%3]
}

// Create roundStart from line of input.
func parseRound(line string) round {
	s := strings.Split(line, " ")
	return round{you: s[1][0], them: s[0][0]}
}

func Part1Solver(in string) int {
	r := 0

	for _, v := range strings.Split(in, "\n") {
		r += parseRound(v).getScore()
	}

	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
