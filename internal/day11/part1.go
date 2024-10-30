package day11

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

type monkey struct {
	items         []int
	op            string
	testDivNumber int
	monkeyTrue    int
	monkeyFalse   int
	inspected     int
}

func parseMonkeys(lines <-chan string) <-chan monkey {
	c := make(chan monkey, 8)
	go func() {
		for v := range lines {
			r := monkey{}
			re := regexp.MustCompile(`Monkey (\d):\n  Starting items: (\d.*)\n  Operation: new = (.*)\n  Test: divisible by (\d.*)\n    If true: throw to monkey (\d.*)\n    If false: throw to monkey (\d.*)`)
			res := re.FindStringSubmatch(v)
			it := []int{}
			for _, v := range strings.Split(res[2], ", ") {
				n, _ := strconv.Atoi(v)
				it = append(it, n)
			}
			r.items = it
			r.op = res[3]
			r.testDivNumber, _ = strconv.Atoi(res[4])
			r.monkeyTrue, _ = strconv.Atoi(res[5])
			r.monkeyFalse, _ = strconv.Atoi(res[6])
			c <- r
		}
		close(c)
	}()
	return c
}

func calculateWorryLevel(item int, op string) int {
	s := strings.Fields(op)
	op1 := 0
	op2 := 0
	switch s[0] {
	case "old":
		op1 = item
	default:
		op1, _ = strconv.Atoi(s[0])
	}
	switch s[2] {
	case "old":
		op2 = item
	default:
		op2, _ = strconv.Atoi(s[2])
	}
	switch s[1] {
	case "+":
		return op1 + op2
	default:
		return op1 * op2
	}
}

func simulateRound(monkeys *[]monkey, worryDecrease int) {
	for i, v := range *monkeys {
		for _, item := range v.items {
			(*monkeys)[i].inspected++
			// Worry level per monkey
			item = calculateWorryLevel(item, v.op)
			// Originaly Divide by 3 after monkey loses interest
			// Since part2 changed the rules we now do mudulo of common multiple
			// This way we avoid overflows of integers but also remain the division properties
			if worryDecrease == 3 {
				item = item / 3
			} else {
				item = item % worryDecrease
			}

			if item%v.testDivNumber == 0 {
				(*monkeys)[v.monkeyTrue].items = append((*monkeys)[v.monkeyTrue].items, item)
			} else {
				(*monkeys)[v.monkeyFalse].items = append((*monkeys)[v.monkeyFalse].items, item)
			}
			(*monkeys)[i].items = (*monkeys)[i].items[1:]
		}
	}
}

func Part1Solver(file string) int {
	monkeys := util.Collect(parseMonkeys(util.ReadParagraphs(file)))
	// Simulate 20 rounds
	for i := 0; i < 20; i++ {
		simulateRound(&monkeys, 3)
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day11/input.txt"))
}
