package day18

import (
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/channels"
)

type coord3 struct {
	X, Y, Z int
}

func parseCoords3(in <-chan string) <-chan coord3 {
	c := make(chan coord3, 50)
	go func() {
		for v := range in {
			s := strings.Split(v, ",")
			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			z, _ := strconv.Atoi(s[2])
			c <- coord3{x, y, z}
		}
		close(c)
	}()
	return c
}

func freeSides(m map[coord3]bool, k coord3) int {
	r := 0
	if _, ok := m[coord3{k.X + 1, k.Y, k.Z}]; !ok {
		r++
	}
	if _, ok := m[coord3{k.X - 1, k.Y, k.Z}]; !ok {
		r++
	}
	if _, ok := m[coord3{k.X, k.Y + 1, k.Z}]; !ok {
		r++
	}
	if _, ok := m[coord3{k.X, k.Y - 1, k.Z}]; !ok {
		r++
	}
	if _, ok := m[coord3{k.X, k.Y, k.Z + 1}]; !ok {
		r++
	}
	if _, ok := m[coord3{k.X, k.Y, k.Z - 1}]; !ok {
		r++
	}

	return r
}

func Part1Solver(file string) int {
	m := make(map[coord3]bool)
	for v := range parseCoords3(channels.ReadLines(file)) {
		m[v] = true
	}

	sides := 0
	for k := range m {
		sides += freeSides(m, k)
	}

	return sides
}
