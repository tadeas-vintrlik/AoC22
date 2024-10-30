package day4

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

// Check if two parts of the clean sections overlap in atleast one part
func (cs cleanSection) overlapSelf() bool {
	return (cs.start2 <= cs.start1 && cs.start1 <= cs.end2) ||
		(cs.start2 <= cs.end1 && cs.end1 <= cs.end2) || cs.fullyContainsSelf()
}

func Part2Solver(file string) int {
	r := 0
	for v := range channels.ReadLines(file) {
		if (parseCleanSection(v)).overlapSelf() {
			r++
		}
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day4/input.txt"))
}
