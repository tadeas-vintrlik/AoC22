package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

type instruction struct {
	name  string
	value int
	delay int
}

func parsePrg(file string) ([]instruction, error) {
	prg := []instruction{}
	for v := range util.ReadLines(file) {
		ins, err := parseInstruction(v)
		if err != nil {
			return []instruction{}, err
		}
		prg = append(prg, ins)
	}
	return prg, nil
}

func parseInstruction(line string) (instruction, error) {
	s := strings.Fields(line)
	ins := instruction{name: s[0]}
	if ins.name == "noop" {
		return ins, nil
	}
	n, err := strconv.Atoi(s[1])
	ins.value = n
	return ins, err
}

func Part1Solver(file string) int {
	prg, err := parsePrg(file)
	if err != nil {
		panic(err)
	}
	cycle := 0
	x := 1
	r := 0
	periodmax := 20
	var delayed *instruction = nil
	ip := 0
	for ip < len(prg) || delayed != nil {
		cycle++

		if cycle == periodmax {
			r += x * cycle
			periodmax += 40
		}

		if delayed == nil {
			ins := prg[ip]
			if ins.name == "addx" {
				ins.delay = 2
				delayed = &ins
			}
			ip++
		}

		if delayed != nil {
			delayed.delay--
			if delayed.delay == 0 {
				x += delayed.value
				delayed = nil
			}
		}
	}
	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day10/input.txt"))
}
