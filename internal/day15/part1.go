package day15

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

type reading struct {
	sensor grid.Coord
	beacon grid.Coord
	dist   int // Manhattan distance from sensor to beacon
}

func parseReading(line string) reading {
	r := reading{}
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	res := re.FindStringSubmatch(line)
	sx, _ := strconv.Atoi(res[1])
	sy, _ := strconv.Atoi(res[2])
	bx, _ := strconv.Atoi(res[3])
	by, _ := strconv.Atoi(res[4])
	r.sensor = grid.Coord{X: sx, Y: sy}
	r.beacon = grid.Coord{X: bx, Y: by}
	r.dist = grid.Distance(r.sensor, r.beacon)
	return r
}

func parseReadings(in <-chan string) <-chan reading {
	c := make(chan reading, 50)
	go func() {
		for v := range in {
			c <- parseReading(v)
		}
		close(c)
	}()
	return c
}

func Part1Solver(file string, line int) int {
	s := channels.Collect(parseReadings(channels.ReadLines(file)))
	xs := slice.Flatten(slice.Map(s, func(r reading) []int {
		return []int{r.sensor.X + r.dist, r.sensor.X - r.dist}
	}))
	xmin := slice.Min(xs)
	xmax := slice.Max(xs)

	y := line
	ret := 0
	for x := xmin; x < xmax; x++ {
		for _, v := range s {
			c := grid.Coord{X: x, Y: y}
			distLine := grid.Distance(c, v.sensor)
			if distLine <= v.dist && !(c == v.beacon || c == v.sensor) {
				ret++
				break
			}
		}
	}

	return ret
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day15/input.txt", 2000000))
}
