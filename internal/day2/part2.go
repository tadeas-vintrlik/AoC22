package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

// X => A  B  C
// Y => C  A  B
// Z => B  C  A
// Similar to round_result in part 1 except now we want value of move
var choice_value = []int{3, 1, 2}

func (r round) getScorePart2() int {
	// value for round outcome is dependent on your character
	// X => 0, Y => 3, Z => 6
	return int(r.you-'X')*3 + choice_value[(int((r.you-'X'))+int((r.them-'A')))%3]
}

func Part2Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		r += parseRound(v).getScorePart2()
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver(input))
}
