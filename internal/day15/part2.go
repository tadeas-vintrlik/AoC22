package day15

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func Part2Solver(file string, max int) int {
	s := channels.Collect(parseReadings(channels.ReadLines(file)))
	for y := 0; y < max; y++ {
		for x := 0; x < max; x++ {
			found := true
			for _, v := range s {
				c := grid.Coord{X: x, Y: y}
				distLine := grid.Distance(c, v.sensor)
				// We are inside range of one of the sensors, here the beacon can't be
				// There we move to the outter edge of the sensor range on the same line
				if distLine <= v.dist {
					x = v.sensor.X + (v.dist - util.Abs(y-v.sensor.Y))
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

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day15/input.txt", 4000000))
}
