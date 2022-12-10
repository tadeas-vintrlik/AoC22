package day1

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

// Find the three elves with most calories in their backpack. Return the sum.
func Part2Solver(file string) int {
	all := []int{}
	c := 0
	for v := range util.ReadLines(file) {
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
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day1/input.txt"))
}
