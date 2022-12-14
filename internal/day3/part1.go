package day3

import (
	"fmt"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

func itemToPriority(b byte) int {
	if b < 'a' {
		return int(b - 'A' + 27)
	}
	return int(b - 'a' + 1)
}

// Find the item shared by both compartments (letter in both halves of the line)
func findSharedItem(l string) byte {
	mid := len(l) / 2
	l1 := []byte(l[:mid])
	l2 := l[mid:]
	for _, v := range l1 {
		if strings.IndexByte(l2, v) != -1 {
			return v
		}
	}
	panic("No shared item found")
}

func Part1Solver(file string) int {
	r := 0
	for v := range channels.ReadLines(file) {
		r += itemToPriority(findSharedItem(v))
	}
	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day3/input.txt"))
}
