package day6

import "fmt"

func Part2Solver(in string) int {
	return bothPartsSolver(in, 14)
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day6/input.txt"))
}
