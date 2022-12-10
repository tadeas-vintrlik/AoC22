package day10

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type instruction struct {
	name  string
	value int
	delay int
}

func parsePrg(in string) ([]instruction, error) {
	prg := make([]instruction, strings.Count(in, "\n")+1)
	for i, v := range strings.Split(in, "\n") {
		ins, err := parseInstruction(v)
		if err != nil {
			return []instruction{}, err
		}
		prg[i] = ins
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

func Part1Solver(in string) int {
	prg, err := parsePrg(in)
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
