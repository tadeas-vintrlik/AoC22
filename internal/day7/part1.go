package day7

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type fsItem struct {
	name     string
	size     int
	children []*fsItem
	parent   *fsItem
}

func (fi fsItem) isDir() bool {
	// In theory this would return false for empty directories
	// But these have size 0 anyways so we do not care
	return len(fi.children) != 0
}

// Parse a ouptut of the shell into a file system. Return the root directory.
// Assumes the input will be correct format.
func parseShell(in string) (fsItem, error) {
	root := fsItem{name: "/", parent: nil}
	var c *fsItem = &root
	for _, v := range strings.Split(in, "\n") {
		s := strings.Fields(v)
		switch s[0] {
		case "$":
			switch s[1] {
			case "cd":
				switch s[2] {
				case "..":
					c = c.parent
				case "/":
					c = &root
				default:
					found := false
					for _, v := range c.children {
						if v.name == s[2] {
							c = v
							found = true
							break
						}
					}
					if !found {
						return root, fmt.Errorf("directory %s contains no such sub-directory %s", c.name, s[2])
					}
				}
			default:
				// List no need to do nothing
			}
		case "dir":
			c.children = append(c.children, &fsItem{name: s[1], parent: c})
		default:
			size, err := strconv.Atoi(s[0])
			if err != nil {
				return root, err
			}
			c.children = append(c.children, &fsItem{name: s[1], size: size})
			// File s[0] is size s[1] is name
		}
	}
	return root, nil
}

func (fi *fsItem) directorySize() int {
	if !fi.isDir() {
		return fi.size
	}
	s := 0
	for _, v := range fi.children {
		s += v.directorySize()
	}
	return s
}

// Change dir size to be sum of all of it's recursive files
func (fi *fsItem) sumDirs() {
	if !fi.isDir() {
		return
	}
	s := 0
	for _, v := range fi.children {
		v.sumDirs()
		v.size = v.directorySize()
		s += v.size
	}
	fi.size = s
}

// Sum directories with size up to 100000
func Part1Value(fi *fsItem) int {
	s := 0
	if !fi.isDir() {
		return s
	} else {
		if fi.size <= 100000 {
			s += fi.size
		}
	}
	for _, v := range fi.children {
		s += Part1Value(v)
	}

	return s
}

func Part1Solver(in string) int {
	root, err := parseShell(in)
	if err != nil {
		panic(err)
	}
	root.sumDirs()
	return Part1Value(&root)
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
