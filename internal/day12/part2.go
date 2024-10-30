package day12

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func Part2Solver(file string) int {
	g := grid.Parse(util.Collect(util.ReadLines(file)), func(c rune) rune { return c })
	start := append(g.Find('a'), g.Find('S')[0])
	end := g.Find('E')[0]
	return util.SliceMin(util.SliceMap(start, func(s grid.Node[rune]) int {
		return bothPartSolver(g, s, end)
	}))
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day12/input.txt"))
}
