package day6

import (
	_ "embed"
)

//go:embed input.txt
var input string

func unique(a string) bool {
	max := len(a)
	for i := 0; i < max; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return false
			}
		}
	}
	return true
}

func Part1Solver(in string) int {
	max := len(in)
	for i := 0; i+4 < max; i++ {
		if unique(in[i : i+4]) {
			return i + 4
		}
	}
	panic("solution not found")
}
