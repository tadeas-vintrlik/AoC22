package day10

import (
	"fmt"
	"strings"
)

type screen struct {
	content [240]bool
}

func (s screen) String() string {
	r := strings.Builder{}
	for i := 0; i <= 200; i += 40 {
		for _, v := range s.content[i : i+40] {
			if v {
				r.WriteByte('#')
			} else {
				r.WriteByte('.')
			}
		}
		r.WriteByte('\n')
	}

	str := r.String()
	return str[:len(str)-1]
}

func Part2Solver(file string) string {
	prg, err := parsePrg(file)
	if err != nil {
		panic(err)
	}
	screen := screen{}
	cycle := 0
	x := 1
	var delayed *instruction = nil
	ip := 0
	for ip < len(prg) || delayed != nil {
		if delayed == nil {
			ins := prg[ip]
			if ins.name == "addx" {
				ins.delay = 2
				delayed = &ins
			}
			ip++
		}

		pi := cycle % 40 // pixel index in row
		if pi == x || pi == x+1 || pi == x-1 {
			screen.content[cycle] = true
		}

		if delayed != nil {
			delayed.delay--
			if delayed.delay == 0 {
				x += delayed.value
				delayed = nil
			}
		}
		cycle++
	}
	return screen.String()
}

func Part2() string {
	return fmt.Sprintf("Part 2:\n%s", Part2Solver("../../internal/day10/input.txt"))
}
