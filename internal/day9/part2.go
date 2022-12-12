package day9

import "fmt"

func Part2Solver(file string) int {
	return bothPartSolver(file, 10)
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day9/input.txt"))
}
