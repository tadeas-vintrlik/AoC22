package day5

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

//go:embed input.txt
var input string

// The crate layout is a map of byte slices. The index is the column number.
// The top of the stack is the last element of each slice.
type crates map[int][]byte

func (c crates) move(ins instruction) {
	from := c[ins.from]
	to := c[ins.to]
	for i := 0; i < ins.amount; i++ {
		to = append(to, from[len(from)-1])
		from = from[:len(from)-1]
	}
	c[ins.from] = from
	c[ins.to] = to
}

func parseCrates(in string) crates {
	ret := make(map[int][]byte)
	for _, v := range strings.Split(in, "\n") {
		if v[1] == '1' {
			break
		}
		for i := 1; i < len(v); i += 4 {
			if v[i] == ' ' {
				continue
			}
			// For index to match the column number
			mi := ((i - 1) / 4) + 1
			ret[mi] = append(ret[mi], v[i])
		}
	}
	for k := range ret {
		util.Reverse(ret[k])
	}
	return ret
}

type instruction struct {
	from, to, amount int
}

func parseInstruction(l string) (instruction, error) {
	var r instruction
	_, err := fmt.Sscanf(l, "move %d from %d to %d", &r.amount, &r.from, &r.to)
	return r, err
}

func Part1Solver(in string) string {
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
		c.move(i)
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

func Part1() string {
	return fmt.Sprintf("Part 1 solution: %s", Part1Solver(input))
}