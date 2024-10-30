package day3

import (
	"fmt"
	"strings"
)

// Find badge item - one byte shared between 3 elves (lines)
func findBadgeItem(l3 [3]string) byte {
	for i := 0; i < len(l3[0]); i++ {
		for j := 0; j < len(l3[1]); j++ {
			for k := 0; k < len(l3[2]); k++ {
				if (l3[0][i] == l3[1][j]) && (l3[1][j] == l3[2][k]) {
					return l3[0][i]
				}
			}
		}
	}
	panic("no badge found")
}

func Part2Solver(in string) int {
	l3 := [3]string{}
	r := 0
	for i, v := range strings.Split(in, "\n") {
		// We want to find badge for each group of 3 elves
		l3[i%3] = v
		if i%3 == 2 {
			r += itemToPriority(findBadgeItem(l3))
		}
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part2 solution: %d", Part2Solver(input))
}
