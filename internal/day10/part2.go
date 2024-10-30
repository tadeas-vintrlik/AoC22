package day10

type screen struct {
	content [240]bool
}

func (s screen) String() string {
	r := []string{}
	for i := 0; i <= 200; i += 40 {
		l := make([]byte, 40)
		for i, v := range s.content[i : i+40] {
			if v {
				l[i] = '#'
			} else {
				l[i] = '.'
			}
		}
		r = append(r, string(l))
	}

	ret := ""
	for _, v := range r {
		ret += v + "\n"
	}
	return ret[:len(ret)-1]
}

func Part2Solver(in string) string {
	prg, err := parsePrg(in)
	if err != nil {
		panic(err)
	}
	screen := screen{}
	cycle := 0
	x := 1
	var delayed *instruction = nil
	ip := 0
	for ip < len(prg) || delayed != nil {
		if delayed == nil {
			ins := prg[ip]
			if ins.name == "addx" {
				ins.delay = 2
				delayed = &ins
			}
			ip++
		}

		pi := cycle % 40 // pixel index in row
		if pi == x || pi == x+1 || pi == x-1 {
			screen.content[cycle] = true
		}

		if delayed != nil {
			delayed.delay--
			if delayed.delay == 0 {
				x += delayed.value
				delayed = nil
			}
		}
		cycle++
	}
	return screen.String()
}
