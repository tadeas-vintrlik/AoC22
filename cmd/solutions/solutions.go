package main

import (
	"fmt"

	"github.com/tadeas-vintrlik/AoC22/internal/day1"
	"github.com/tadeas-vintrlik/AoC22/internal/day2"
	"github.com/tadeas-vintrlik/AoC22/internal/day3"
	"github.com/tadeas-vintrlik/AoC22/internal/day4"
	"github.com/tadeas-vintrlik/AoC22/internal/day5"
	"github.com/tadeas-vintrlik/AoC22/internal/day6"
)

type partFunc func() string

type solFunc [2]partFunc

var s []solFunc = []solFunc{
	{day1.Part1, day1.Part2},
	{day2.Part1, day2.Part2},
	{day3.Part1, day3.Part2},
	{day4.Part1, day4.Part2},
	{day5.Part1, day5.Part2},
	{day6.Part1, day6.Part2},
}

func main() {
	for i, v := range s {
		fmt.Printf("Day %d: %s, %s\n", i+1, v[0](), v[1]())
	}
}