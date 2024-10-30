package day1

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Part2Solver(in string) int {
	all := []int{}
	current := 0
	for _, v := range strings.Split(in, "\n") {
		if v == "" {
			all = append(all, current)
			current = 0
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		current += i
	}

	all = append(all, current)

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
