package day2

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Reponse of you/oponent Rock/paper/scissors
type resp byte

// o prefix means oponent
var (
	rock      resp = 'X'
	paper     resp = 'Y'
	scissors  resp = 'Z'
	orock     resp = 'A'
	opaper    resp = 'B'
	oscissors resp = 'C'
)

type outcome int

var (
	defeat  outcome = 0
	tie     outcome = 1
	victory outcome = 2
)

// Score depedening on outcome (used as index)
var score = []int{0, 3, 6}

func (r resp) unexpected() string {
	return fmt.Sprintf("invalid: %c", r)
}

// Expects a valid resp. Return score for resp type.
func (r resp) getScore() int {
	switch r {
	case 'X':
		return 1
	case 'Y':
		return 2
	case 'Z':
		return 3
	default:
		panic(r.unexpected())
	}
}

// If r beats o.
func (r resp) beats(o resp) bool {
	return (o == orock && r == paper) || (o == opaper && r == scissors) || (o == oscissors && r == rock)
}

// If r ties with o.
func (r resp) ties(o resp) bool {
	return (o == orock && r == rock) || (o == opaper && r == paper) || (o == oscissors && r == scissors)
}

// Structure containing one line of input (response of you and your oponent)
type roundStrat struct {
	oponent resp
	you     resp
}

// Get score for the outcome of the round.
func (rs roundStrat) getRoundOutcome() outcome {
	if rs.you.ties(rs.oponent) {
		return tie
	}
	if rs.you.beats(rs.oponent) {
		return victory
	}
	return defeat
}

// Get score for both outcome and your response.
func (rs roundStrat) getTotalScore() int {
	return score[rs.getRoundOutcome()] + rs.you.getScore()
}

// Create roundStart from line of input.
func parseRoundStrat(line string) roundStrat {
	s := strings.Split(line, " ")
	o := s[0][0]
	y := s[1][0]
	you := resp(y)
	them := resp(o)
	return roundStrat{you: you, oponent: them}
}

func Part1Solver(in string) int {
	r := 0

	for _, v := range strings.Split(in, "\n") {
		r += parseRoundStrat(v).getTotalScore()
	}

	return r
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver(input))
}
