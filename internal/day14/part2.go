package day14

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

func Part2Solver(file string) int {
	paths := channels.Collect(parseRockPaths(channels.ReadLines(file)))
	allCoords := slice.Flatten(slice.Map(paths, func(p grid.Path) []grid.Coord {
		return slice.Map(p, func(c grid.Coord) grid.Coord {
			return c
		})
	}))
	// We extend the bounds a little to fit for part2
	xmax := slice.Max(slice.Map(allCoords, func(c grid.Coord) int {
		return c.X
	})) * 2
	ymax := slice.Max(slice.Map(allCoords, func(c grid.Coord) int {
		return c.Y
	})) + 2
	g := grid.New[rune](xmax+1, ymax+1)
	g.Fill('.')
	paths = append(paths, grid.Path{grid.Coord{X: 0, Y: ymax}, grid.Coord{X: xmax, Y: ymax}})
	for _, v := range paths {
		g.FillPath(v, '#')
	}

	sand := 0
	done := false
	for !done {
		x := 500
		y := 0

		for {
			if g.Get(x, y+1) == '.' {
				y++
			} else if g.Get(x-1, y+1) == '.' {
				y++
				x--
			} else if g.Get(x+1, y+1) == '.' {
				y++
				x++
			} else {
				// Stabilised
				g.Set(x, y, 'o')
				sand++
				if x == 500 && y == 0 {
					done = true
				}
				break
			}
		}
	}
	return sand
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day14/input.txt"))
}
