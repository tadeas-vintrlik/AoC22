package day15

import (
	"regexp"
	"strconv"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/grid"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

type reading struct {
	sensor grid.Coord
	beacon grid.Coord
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
		d := grid.Distance(r.beacon, r.sensor)
		return []int{r.sensor.X + d, r.sensor.X - d}
	}))
	ys := slice.Flatten(slice.Map(s, func(r reading) []int {
		d := grid.Distance(r.beacon, r.sensor)
		return []int{r.sensor.Y + d, r.sensor.Y - d}
	}))
	xmin := slice.Min(xs)
	xmax := slice.Max(xs)
	ymin := slice.Min(ys)
	ymax := slice.Max(ys)

	g := grid.NewNonZero[string](xmin, xmax, ymin, ymax)
	g.Fill(".")
	for _, v := range s {
		g.FillAround(v.sensor, grid.Distance(v.sensor, v.beacon), "#")
	}
	for _, v := range s {
		g.Set(v.sensor.X, v.sensor.Y, "S")
		g.Set(v.beacon.X, v.beacon.Y, "B")
	}

	y := line
	ret := 0
	for x := g.MinX(); x < g.MaxX(); x++ {
		if g.Get(x, y) == "#" {
			ret++
		}
	}

	return ret
}
