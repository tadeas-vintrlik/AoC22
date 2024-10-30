package day7

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/pkg/tree"
)

func Part2Solver(in string) int {
	root, err := parseShell(in)
	if err != nil {
		panic(err)
	}
	sumDirs(root)

	total := 70000000
	desired := 30000000
	occupied := root.Value.size

	min := root.Value.size
	tree.PostorderCallback(root, func(t *tree.Tree[fsItem]) {
		if t.Value.isDir && total-occupied+t.Value.size >= desired && t.Value.size < min {
			min = t.Value.size
		}
	})

	return min
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day7/input.txt"))
}
