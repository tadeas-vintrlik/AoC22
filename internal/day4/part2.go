package day4

import (
	"fmt"
	"strings"
)

// Check if two parts of the clean sections overlap in atleast one part
func (cs cleanSection) overlapSelf() bool {
	return (cs.start2 <= cs.start1 && cs.start1 <= cs.end2) ||
		(cs.start2 <= cs.end1 && cs.end1 <= cs.end2) || cs.fullyContainsSelf()
}

func Part2Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		if (parseCleanSection(v)).overlapSelf() {
			r++
		}
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver(input))
}
