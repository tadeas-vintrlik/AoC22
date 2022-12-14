package day3

import (
	"fmt"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

// Find badge item - one byte shared between 3 elves (lines)
func findBadgeItem(l3 [3]string) byte {
	for i := 0; i < len(l3[0]); i++ {
		if strings.IndexByte(l3[1], l3[0][i]) != -1 &&
			strings.IndexByte(l3[2], l3[0][i]) != -1 {
			return l3[0][i]
		}
	}
	panic("no badge found")
}

func Part2Solver(file string) int {
	l3 := [3]string{}
	r := 0
	i := 0
	for v := range channels.ReadLines(file) {
		// We want to find badge for each group of 3 elves
		l3[i%3] = v
		if i%3 == 2 {
			r += itemToPriority(findBadgeItem(l3))
		}
		i++
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day3/input.txt"))
}
