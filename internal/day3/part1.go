package day3

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func itemToPriority(b byte) int {
	if b < 'a' {
		return int(b - 'A' + 27)
	}
	return int(b - 'a' + 1)
}

// Find the item shared by both compartments (letter in both halves of the line)
func findSharedItem(l string) byte {
	max := len(l)
	mid := max / 2
	for i := 0; i < mid; i++ {
		for j := mid; j < max; j++ {
			if l[i] == l[j] {
				return l[i]
			}
		}
	}
	panic("No shared item found")
}

func Part1Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		r += itemToPriority(findSharedItem(v))

	}
	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1 result: %d", Part1Solver(input))
}
