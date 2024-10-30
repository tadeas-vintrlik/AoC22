package day12

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func climable(c grid.Node[rune], s []grid.Node[rune]) []grid.Node[rune] {
	var ret []grid.Node[rune]
	src := c.Value
	if src == 'S' {
		src = 'a'
	}
	for _, v := range s {
		dest := v.Value
		if dest == 'E' {
			dest = 'z'
		}
		dc := dest - src
		if dc <= 0 || dc == 1 {
			ret = append(ret, v)
		}
	}
	return ret
}

func bothPartSolver(g grid.Grid[rune], start, end grid.Node[rune]) int {
	dist := make(map[grid.Node[rune]]int)
	dist[start] = 0
	g.BFS(
		start, // Root of BFS
		func(from grid.Node[rune]) []grid.Node[rune] {
			return climable(from, g.GetNeigbhours(from.X, from.Y))
		}, // Neighbours callback to get neighbouring nodes
		func(parent, child grid.Node[rune]) {
			dist[child] = dist[parent] + 1
		}, // Callback on parent and child in the BFS order
	)
	if v, ok := dist[end]; ok {
		return v
	} else {
		// In case end is unreachable return maximum distance
		return g.SizeX() * g.SizeY()
	}
}

func Part1Solver(file string) int {
	g := grid.Parse(util.Collect(util.ReadLines(file)), func(c rune) rune { return c })
	start := g.Find('S')[0]
	end := g.Find('E')[0]
	return bothPartSolver(g, start, end)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day12/input.txt"))
}
