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

// TODO: This would be solved with a generic graph structure and BFS implement in pkg

func findMinDistance(g grid.Grid[rune], root grid.Node[rune]) int {
	dist := make(map[grid.Node[rune]]int)
	visited := []grid.Node[rune]{root}
	queue := []grid.Node[rune]{root}
	dist[root] = 0

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		if c.Value == 'E' {
			return dist[c]
		}
		for _, v := range climable(c, g.GetNeigbhours(c.X, c.Y)) {
			if util.SliceContains(visited, v) {
				continue
			}
			dist[v] = dist[c] + 1
			visited = append(visited, v)
			queue = append(queue, v)
		}
	}

	return g.SizeX() * g.SizeY()
}

func Part1Solver(file string) int {
	g := grid.Parse(util.Collect(util.ReadLines(file)), func(c rune) rune { return c })
	start := g.Find('S')
	if len(start) != 1 {
		panic("Did not found starting point")
	}
	s := start[0]
	return findMinDistance(g, s)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day12/input.txt"))
}
