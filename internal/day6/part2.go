package day6

func Part2Solver(in string) int {
	max := len(in)
	for i := 0; i+14 < max; i++ {
		if unique(in[i : i+14]) {
			return i + 14
		}
	}
	panic("solution not found")
}
