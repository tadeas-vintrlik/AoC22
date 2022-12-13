package day13

import (
	"fmt"
	"sort"

	"github.com/tadeas-vintrlik/AoC22/pkg/util"
)

func Part2Solver(file string) int {
	divPackets := []packet{
		{containsList: true, list: []*packet{
			{containsList: true, list: []*packet{
				{val: 2},
			}},
		}},
		{containsList: true, list: []*packet{
			{containsList: true, list: []*packet{
				{val: 6},
			}},
		}},
	}
	packets := util.Collect(util.Flatten(parsePacketPairs(util.ReadParagraphs(file))))
	packets = append(packets, divPackets...)
	sort.SliceStable(packets, func(i, j int) bool {
		return correctOrder([]packet{packets[i], packets[j]}) == right
	})
	ret := 1
	for i, v := range packets {
		if v.containsList && len(v.list) == 1 && v.list[0].containsList && len(v.list[0].list) == 1 {
			if v.list[0].list[0].val == 2 || v.list[0].list[0].val == 6 {
				ret *= i + 1
			}
		}
	}
	return ret
}

func Part2() string {
	return fmt.Sprintf("Part 2: %d", Part2Solver("../../internal/day13/input.txt"))
}
