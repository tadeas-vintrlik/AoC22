package day18

import (
	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

func Part2Solver(file string) int {
	s := channels.Collect(parseCoords3(channels.ReadLines(file)))
	m := make(map[coord3]bool)
	for _, v := range s {
		m[v] = true
	}

	sides := 0

	// Found a bounding box
	xmax := slice.Max(slice.Map(s, func(c coord3) int { return c.X + 1 }))
	ymax := slice.Max(slice.Map(s, func(c coord3) int { return c.Y + 1 }))
	zmax := slice.Max(slice.Map(s, func(c coord3) int { return c.Z + 1 }))

	// Flood fill
	start := coord3{xmax, ymax, zmax}
	q := []coord3{start}
	seen := make(map[coord3]bool)
	for len(q) != 0 {
		c := q[0]
		q = q[1:]

		_, filled := seen[c]
		if filled || c.X > xmax || c.X < -1 || c.Y > ymax || c.Y < -1 || c.Z > zmax || c.Z < -1 {
			continue
		}

		if _, ok := m[c]; ok {
			sides++
			continue
		}

		seen[c] = true
		q = append(q, coord3{c.X + 1, c.Y, c.Z})
		q = append(q, coord3{c.X - 1, c.Y, c.Z})
		q = append(q, coord3{c.X, c.Y + 1, c.Z})
		q = append(q, coord3{c.X, c.Y - 1, c.Z})
		q = append(q, coord3{c.X, c.Y, c.Z + 1})
		q = append(q, coord3{c.X, c.Y, c.Z - 1})
	}

	return sides
}
