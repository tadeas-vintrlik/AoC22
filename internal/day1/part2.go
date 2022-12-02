package day1

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Find the three elves with most calories in their backpack. Return the sum.
func Part2Solver(in string) int {
	all := []int{}
	c := 0
	for _, v := range strings.Split(in, "\n") {
		if v == "" {
			all = append(all, c)
			c = 0
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		c += i
	}

	// Last one will not be appended
	all = append(all, c)

	sort.Ints(all)
	sum := 0
	for _, v := range all[len(all)-3:] {
		sum += v
	}

	return sum
}

func Part2() string {
	return fmt.Sprintf("Part 2 solution: %d", Part2Solver((input)))
}
