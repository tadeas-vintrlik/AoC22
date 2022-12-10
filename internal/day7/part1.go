package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/tree"
	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

type fsItem struct {
	name  string
	size  int
	isDir bool
}

// Parse a ouptut of the shell into a file system. Return the root directory.
// Assumes the input will be correct format.
func parseShell(file string) (*tree.Tree[fsItem], error) {
	root := tree.Tree[fsItem]{Value: fsItem{name: "/", isDir: true}}
	var c *tree.Tree[fsItem] = &root
	for v := range util.ReadLines(file) {
		s := strings.Fields(v)
		switch s[0] {
		case "$":
			switch s[1] {
			case "cd":
				switch s[2] {
				case "..":
					c = tree.Parent(c)
				case "/":
					c = &root
				default:
					found := false
					for _, v := range tree.Children(c) {
						if v.Value.name == s[2] {
							c = v
							found = true
							break
						}
					}
					if !found {
						return &root, fmt.Errorf("directory %s contains no such sub-directory %s", c.Value.name, s[2])
					}
				}
			default:
				// List no need to do nothing
			}
		case "dir":
			tree.Append(c, fsItem{name: s[1], isDir: true})
		default:
			size, err := strconv.Atoi(s[0])
			if err != nil {
				return &root, err
			}
			tree.Append(c, fsItem{name: s[1], size: size})
			// File s[0] is size s[1] is name
		}
	}
	return &root, nil
}

// Change dir size to be sum of all of it's recursive files
func sumDirs(ft *tree.Tree[fsItem]) {
	tree.PostorderCallback(ft, func(n *tree.Tree[fsItem]) {
		for _, v := range tree.Children(n) {
			n.Value.size += v.Value.size
		}
	})
}

func Part1Solver(file string) int {
	root, err := parseShell(file)
	if err != nil {
		panic(err)
	}
	sumDirs(root)

	// Sum directories with size up to 100000
	s := 0
	tree.PostorderCallback(root, func(t *tree.Tree[fsItem]) {
		if t.Value.isDir && t.Value.size <= 100000 {
			s += t.Value.size
		}
	})
	return s
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day7/input.txt"))
}
