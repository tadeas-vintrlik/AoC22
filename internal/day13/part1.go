package day13

import (
	"fmt"
	"go/scanner"
	"go/token"
	"strconv"
	"strings"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

type packet struct {
	val          int
	containsList bool
	list         []*packet
	parent       *packet
}

func (p packet) isList() bool {
	return p.containsList
}

// Could have parsed it as a JSON array, that would however require any type and type casting.
// This way we create a generic int list parser that is type safe.
func parsePacket(line string) packet {
	ret := packet{parent: nil, containsList: true}
	current := &ret
	inb := []byte(line)
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(line))
	s.Init(file, inb, nil, scanner.ScanComments)
	end := false
	for {
		if end {
			break
		}
		_, tok, lit := s.Scan()
		switch tok {
		case token.INT:
			val, err := strconv.Atoi(lit)
			if err != nil {
				panic(err)
			}
			n := &packet{val: val}
			current.list = append(current.list, n)
		case token.LBRACK:
			n := &packet{parent: current, containsList: true}
			current.list = append(current.list, n)
			current = n
		case token.RBRACK:
			current = current.parent
		case token.EOF:
			end = true
		}
	}
	return ret
}

func parsePacketPairs(in <-chan string) <-chan []packet {
	c := make(chan []packet, 50)
	go func() {
		for v := range in {
			s := strings.Split(v, "\n")
			if len(s) != 3 && s[2] != "" {
				panic("Expected 2 lines per packet pair")
			}
			packets := make([]packet, 2)
			for i := 0; i < 2; i++ {
				packets[i] = parsePacket(s[i])
			}
			c <- packets
		}
		close(c)
	}()
	return c
}

type compared int

const (
	right compared = 1
	wrong compared = 0
	next  compared = -1 // Can't decide, continue
)

func correctOrder(in []packet) compared {
	c1 := in[0]
	c2 := in[1]
	if !c1.isList() && !c2.isList() {
		if c1.val < c2.val {
			return right
		}
		if c1.val > c2.val {
			return wrong
		}
		if c1.val == c2.val {
			return next
		}
	}
	if c1.isList() && c2.isList() {
		maxi := util.Min([]int{len(c1.list), len(c2.list)})
		for i := 0; i < maxi; i++ {
			r := correctOrder([]packet{*c1.list[i], *c2.list[i]})
			if r == wrong {
				return wrong
			}
			if r == right {
				return right
			}
		}
		if len(c1.list) < len(c2.list) {
			return right
		}
		if len(c1.list) > len(c2.list) {
			return wrong
		}
		return next
	}
	if c1.isList() && !c2.isList() {
		return correctOrder([]packet{c1, {containsList: true, list: []*packet{&c2}}})
	}
	if !c1.isList() && c2.isList() {
		return correctOrder([]packet{{containsList: true, list: []*packet{&c1}}, c2})
	}

	panic("Should not be reachable with valid input")
}

func checkOrders(in <-chan []packet) <-chan int {
	c := make(chan int)
	i := 0
	go func() {
		for v := range in {
			i++
			r := correctOrder(v)
			if r == next {
				panic("Should not happen")
			}
			if r == right {
				c <- i
			}
		}
		close(c)
	}()
	return c
}

func Part1Solver(file string) int {
	return util.Sum(checkOrders(parsePacketPairs(util.ReadParagraphs(file))))
}

func Part1() string {
	return fmt.Sprintf("Part 1: %d", Part1Solver("../../internal/day13/input.txt"))
}
