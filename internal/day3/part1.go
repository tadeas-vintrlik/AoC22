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

func Part1Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		r += itemToPriority(findSharedItem(v))
	}
	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
