package day4

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type cleanSection struct {
	start1, end1, start2, end2 int
}

// Check if one part of cleanSection fully contains the other
func (cs cleanSection) fullyContainsSelf() bool {
	return (cs.start1 >= cs.start2 && cs.end1 <= cs.end2) ||
		(cs.start2 >= cs.start1 && cs.end2 <= cs.end1)
}

func parseCleanSection(line string) cleanSection {
	s := strings.Split(line, ",")
	s1 := strings.Split(s[0], "-")
	s2 := strings.Split(s[1], "-")
	start1, _ := strconv.Atoi(s1[0])
	end1, _ := strconv.Atoi(s1[1])
	start2, _ := strconv.Atoi(s2[0])
	end2, _ := strconv.Atoi(s2[1])
	return cleanSection{start1, end1, start2, end2}
}

func Part1Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		if (parseCleanSection(v)).fullyContainsSelf() {
			r++
		}
	}
	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
