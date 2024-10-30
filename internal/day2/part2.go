package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

// Get my response depending on o (oponent) response and the desired outcome.
func (o resp) getMyResp(oc outcome) resp {
	switch o {
	case orock:
		return [3]resp{scissors, rock, paper}[oc]
	case opaper:
		return [3]resp{rock, paper, scissors}[oc]
	case oscissors:
		return [3]resp{paper, scissors, rock}[oc]
	default:
		panic(o.unexpected())
	}
}

// Part2 has new meaning for these characters these are infact desired outcomes.
func (o resp) getRoundOutcomePart2() outcome {
	switch o {
	case 'X':
		return defeat
	case 'Y':
		return tie
	case 'Z':
		return victory
	default:
		panic(o.unexpected())
	}
}

// Get score for both the round outcome and our response.
func (rs roundStrat) getTotalScorePart2() int {
	oc := rs.you.getRoundOutcomePart2()
	r := rs.oponent.getMyResp(oc)
	return r.getScore() + score[oc]
}

func Part2Solver(in string) int {
	r := 0
	for _, v := range strings.Split(in, "\n") {
		r += parseRoundStrat(v).getTotalScorePart2()
	}
	return r
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver(input))
}
