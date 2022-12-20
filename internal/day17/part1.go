package day17

import (
	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

var shapes = [][]grid.Coord{
	{{X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0}, {X: 5, Y: 0}},               //  horizontal line
	{{X: 3, Y: 2}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 3, Y: 0}}, // plus
	{{X: 4, Y: 2}, {X: 4, Y: 1}, {X: 4, Y: 0}, {X: 3, Y: 0}, {X: 2, Y: 0}}, // L shape
	{{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}, {X: 2, Y: 3}},               // vertical line
	{{X: 2, Y: 0}, {X: 3, Y: 0}, {X: 2, Y: 1}, {X: 3, Y: 1}},               // box
}

func Part1Solver(file string) int {
	jet_pattern := channels.Collect(channels.ReadLines(file))[0]
	chamber := []grid.Coord{}
	ytop := 0
	irock := 0
	ijet := 0

	for irock != 2022 {
		// Add a new rock 3 spaces above the topmost
		rock := slice.Map(shapes[irock%len(shapes)], func(c grid.Coord) grid.Coord {
			return grid.Coord{X: c.X, Y: c.Y + 3 + ytop}
		})

		for {
			// See if there are any collisions with already existing rocks, if not move in jet direction
			dx := int(jet_pattern[ijet%len(jet_pattern)]) - int('=') // -1 for <, 1 for > see ASCII table
			ijet++
			xcollisions := slice.Filter(rock, func(c grid.Coord) bool {
				nx := c.X + dx
				if nx > 6 || nx < 0 {
					// Overflow from chamber
					return true
				}

				if slice.Contains(chamber, grid.Coord{X: nx, Y: c.Y}) {
					return true
				}

				return false
			})

			if len(xcollisions) == 0 {
				rock = slice.Map(rock, func(c grid.Coord) grid.Coord {
					return grid.Coord{X: c.X + dx, Y: c.Y}
				})
			}

			ycollisions := slice.Filter(rock, func(c grid.Coord) bool {
				ny := c.Y - 1
				if ny < 0 {
					// Overflow from chamber
					return true
				}

				if slice.Contains(chamber, grid.Coord{X: c.X, Y: ny}) {
					return true
				}

				return false
			})

			if len(ycollisions) == 0 {
				rock = slice.Map(rock, func(c grid.Coord) grid.Coord {
					return grid.Coord{X: c.X, Y: c.Y - 1}
				})
			} else {
				// rock comes to rest
				chamber = append(chamber, rock...)
				irock++
				ytop = slice.Max(slice.Map(chamber, func(c grid.Coord) int {
					return c.Y + 1
				}))
				break
			}
		}
	}

	return ytop
}
