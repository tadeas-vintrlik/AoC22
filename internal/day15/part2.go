package day15

import (
	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
)

func Part2Solver(file string, max int) int {
	s := channels.Collect(parseReadings(channels.ReadLines(file)))
	for y := 0; y < max; y++ {
		for x := 0; x < max; x++ {
			found := true
			for _, v := range s {
				c := grid.Coord{X: x, Y: y}
				distLine := grid.Distance(c, v.sensor)
				distBeacon := grid.Distance(v.beacon, v.sensor)
				if distLine <= distBeacon || c == v.beacon || c == v.sensor {
					found = false
					break
				}
			}
			if found {
				return x*4000000 + y
			}
		}
	}
	panic("Solution not found")
}
