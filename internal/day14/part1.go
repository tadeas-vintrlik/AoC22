package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

func parseRockPaths(lines <-chan string) <-chan grid.Path {
	c := make(chan grid.Path, 50)
	go func() {
		for v := range lines {
			s := strings.Fields(v)
			p := grid.Path{}
			for i, coord := range s {
				if i%2 == 1 {
					continue
				}
				s2 := strings.Split(coord, ",")
				x, err := strconv.Atoi(s2[0])
				if err != nil {
					panic(err)
				}
				y, err := strconv.Atoi(s2[1])
				if err != nil {
					panic(err)
				}
				p = append(p, grid.Coord{X: x, Y: y})
			}
			c <- p
		}
		close(c)
	}()
	return c
}

func Part1Solver(file string) int {
	paths := channels.Collect(parseRockPaths(channels.ReadLines(file)))
	allCoords := slice.Flatten(slice.Map(paths, func(p grid.Path) []grid.Coord {
		return slice.Map(p, func(c grid.Coord) grid.Coord {
			return c
		})
	}))
	xmax := slice.Max(slice.Map(allCoords, func(c grid.Coord) int {
		return c.X
	}))
	ymax := slice.Max(slice.Map(allCoords, func(c grid.Coord) int {
		return c.Y
	}))
	g := grid.New[rune](xmax+1, ymax+1)
	g.Fill('.')
	for _, v := range paths {
		g.FillPath(v, '#')
	}

	sand := 0
	flowOut := false
	for !flowOut {
		x := 500
		y := 0

		for {
			if y+1 == g.SizeY() {
				flowOut = true
				break
			}

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
				break
			}
		}
	}
	return sand
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day14/input.txt"))
}
