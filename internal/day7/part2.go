package day7

import "fmt"

func (fi *fsItem) collectAllDirs(t *[]*fsItem) {
	if !fi.isDir() {
		return
	}
	*t = append(*t, fi)
	for _, v := range fi.children {
		v.collectAllDirs(t)
	}
}

func Part2Solver(in string) int {
	root, err := parseShell(in)
	if err != nil {
		panic(err)
	}
	root.sumDirs()
	var allDirs []*fsItem
	root.collectAllDirs(&allDirs)

	total := 70000000
	desired := 30000000
	occupied := root.size

	min := root.size
	for _, v := range allDirs {
		if total-occupied+v.size >= desired && v.size < min {
			min = v.size
		}
	}
	return min
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver(input))
}
