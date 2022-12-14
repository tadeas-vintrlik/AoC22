package day1

import (
	"fmt"
	"strconv"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

// Find the elf with most calories in their backpack. Return the number of calories
// There is an empty line between each elf.
func Part1Solver(file string) int {
	max := 0
	c := 0
	for v := range channels.ReadLines(file) {
		if v == "" {
			if c > max {
				max = c
			}
			c = 0
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		c += i
	}

	// The last one will not be checked
	if c > max {
		return c
	} else {
		return max
	}
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day1/input.txt"))
}
