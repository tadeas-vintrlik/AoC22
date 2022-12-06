package day6

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func unique(a string) bool {
	max := len(a)
	for i := 0; i < max; i++ {
		if strings.IndexByte(a[i+1:], a[i]) != -1 {
			return false
		}
	}
	return true
}

func bothPartsSolver(in string, size int) int {
	max := len(in)
	for i := 0; i+size < max; i++ {
		if unique(in[i : i+size]) {
			return i + size
		}
	}
	panic("solution not found")
}

func Part1Solver(in string) int {
	return bothPartsSolver(in, 4)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
