package day6

import (
	"fmt"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

func unique(a string) bool {
	max := len(a)
	for i := 0; i < max; i++ {
		if strings.IndexByte(a[i+1:], a[i]) != -1 {
			return false
		}
	}
	return true
}

func bothPartsSolver(file string, size int) int {
	in := <-channels.ReadLines(file)
	max := len(in)
	for i := 0; i+size < max; i++ {
		if unique(in[i : i+size]) {
			return i + size
		}
	}
	panic("solution not found")
}

func Part1Solver(file string) int {
	return bothPartsSolver(file, 4)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day6/input.txt"))
}
