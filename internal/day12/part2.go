package day12

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func Part2Solver(file string) int {
	g := grid.Parse(util.Collect(util.ReadLines(file)), func(c rune) rune { return c })
	start := append(g.Find('a'), g.Find('S')[0])

	c := make(chan int, len(start))
	for _, v := range start {
		go func(v grid.Node[rune]) {
			start := v
			end := g.Find('E')[0]
			c <- bothPartSolver(g, start, end)
		}(v)
	}

	min := g.SizeX() * g.SizeY()
	for i := 0; i < len(start); i++ {
		v := <-c
		if v < min {
			min = v
		}
	}

	return min
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day12/input.txt"))
}
