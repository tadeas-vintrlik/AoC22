package day11

import (
	"fmt"
	"sort"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func Part2Solver(file string) int {
	monkeys := util.Collect(parseMonkeys(util.ReadParagraphs(file)))
	// Find a common multiple
	cm := 1
	for _, v := range monkeys {
		cm *= int(v.testDivNumber)
	}
	// Simulate 10000 rounds
	for i := 0; i < 10000; i++ {
		simulateRound(&monkeys, cm)
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day11/input.txt"))
}
