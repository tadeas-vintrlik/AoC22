package day8

func scenicScore(m [][]int, i, j int) int {
	c := m[i][j]
	st := 0
	for ti := i + 1; ti < len(m); ti++ {
		st++
		if m[ti][j] >= c {
			break
		}
	}

	sb := 0
	for ti := i - 1; ti >= 0; ti-- {
		sb++
		if m[ti][j] >= c {
			break
		}
	}

	sr := 0
	for tj := j + 1; tj < len(m[0]); tj++ {
		sr++
		if m[i][tj] >= c {
			break
		}
	}

	sl := 0
	for tj := j - 1; tj >= 0; tj-- {
		sl++
		if m[i][tj] >= c {
			break
		}
	}

	return sl * sr * st * sb
}

func Part2Solver(in string) int {
	m := parseInput(in)
	maxi := len(m) - 1
	maxj := len(m[0]) - 1

	// Count scenic score
	ss := 0
	for i := 1; i < maxi; i++ {
		for j := 1; j < maxj; j++ {
			if visible(m, i, j) {
				c := scenicScore(m, i, j)
				if c > ss {
					ss = c
				}
			}
		}
	}

	return ss
}
