package day5

import (
	"fmt"
	"strings"
)

// Now can move multiple crates at once
func (c crates) movePart2(ins instruction) {
	from := c[ins.from]
	to := c[ins.to]
	to = append(to, from[(len(from)-ins.amount):]...)
	from = from[:len(from)-ins.amount]
	c[ins.from] = from
	c[ins.to] = to
}

func Part2Solver(in string) string {
	c := parseCrates(in)
	begun := false
	for _, v := range strings.Split(in, "\n") {
		if v == "" {
			begun = true
			continue
		}
		if !begun {
			continue
		}

		i, err := parseInstruction(v)
		if err != nil {
			fmt.Println(v)
			panic(err)
		}
		c.movePart2(i)
	}

	var sb strings.Builder
	for i := 1; ; i++ {
		if _, ok := c[i]; !ok {
			break
		}
		v := c[i]
		sb.WriteByte(v[len(v)-1])
	}
	return sb.String()
}

func Part2() string {
	return fmt.Sprintf("Part 2 solution: %s", Part2Solver(input))
}
