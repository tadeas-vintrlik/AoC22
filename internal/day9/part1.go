package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
	"github.com/tadeas-vintrlik/AoC22/pkg/slice"
)

type coordinate struct {
	x, y int
}

type instruction struct {
	direction byte
	distance  int
}

func moveRope(cs []coordinate, ins instruction) []coordinate {
	head := cs[len(cs)-1]
	switch ins.direction {
	case 'U':
		head.y++
	case 'D':
		head.y--
	case 'L':
		head.x--
	case 'R':
		head.x++
	default:
		panic(fmt.Sprintf("Unexpected movement: %c", ins.direction))
	}
	cs[len(cs)-1] = head
	for i := len(cs) - 2; i >= 0; i-- {
		// C is current part of rope, b is the one before (closer to head or head itself)
		c := &cs[i]
		b := cs[i+1]
		dx := b.x - c.x
		dy := b.y - c.y
		if dx > 1 || dy > 1 || dx < -1 || dy < -1 {
			if dx > 0 {
				c.x++
			} else if dx < 0 {
				c.x--
			}
			if dy > 0 {
				c.y++
			} else if dy < 0 {
				c.y--
			}
		}
	}
	return cs
}

func bothPartSolver(file string, length int) int {
	rope := make([]coordinate, length)
	visited := []coordinate{}
	for v := range channels.ReadLines(file) {
		s := strings.Fields(v)
		d, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		ins := instruction{direction: s[0][0], distance: d}

		for i := 0; i < ins.distance; i++ {
			rope = moveRope(rope, ins)
			if !slice.Contains(visited, rope[0]) {
				visited = append(visited, rope[0])
			}
		}
	}
	return len(visited)
}

func Part1Solver(file string) int {
	return bothPartSolver(file, 2)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day9/input.txt"))
}
