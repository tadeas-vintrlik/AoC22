package day8

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func parseInput(in string) [][]int {
	r := make([][]int, strings.Count(in, "\n")+1)
	for i, v := range strings.Split(in, "\n") {
		r[i] = make([]int, len(v))
		for j, b := range v {
			r[i][j] = int(b - '0')
		}
	}
	return r
}

func visible(m [][]int, i, j int) bool {
	c := m[i][j]
	visible := true
	for ti := i + 1; ti < len(m); ti++ {
		if m[ti][j] >= c {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true
	for ti := i - 1; ti >= 0; ti-- {
		if m[ti][j] >= c {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true
	for tj := j + 1; tj < len(m[0]); tj++ {
		if m[i][tj] >= c {
			visible = false
			break
		}
	}
	if visible {
		return true
	}
	visible = true
	for tj := j - 1; tj >= 0; tj-- {
		if m[i][tj] >= c {
			visible = false
			break
		}
	}
	return visible
}

func Part1Solver(in string) int {
	m := parseInput(in)
	maxi := len(m) - 1
	maxj := len(m[0]) - 1

	// Trees on the border
	v := ((maxi + 1) * (maxi + 1)) - ((maxj - 1) * (maxj - 1))
	for i := 1; i < maxi; i++ {
		for j := 1; j < maxj; j++ {
			if visible(m, i, j) {
				v++
			}
		}
	}

	return v
}
