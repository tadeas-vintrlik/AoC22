package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Part1Solver(in string) int {
	max := 0
	current := 0
	for _, v := range strings.Split(in, "\n") {
		if v == "" {
			if current > max {
				max = current
			}
			current = 0
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		current += i
	}

	// The last one might not be checked
	if current > max {
		return current
	} else {
		return max
	}
}

func Part1() string {
	return fmt.Sprintf("Part 1 solution: %d", Part1Solver((input)))
}
