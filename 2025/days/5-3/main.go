package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

type qualityCalc struct {
	id      int
	lines   [][3]int
	quality int
}

func main() {
	rows, err := util.ReadStrings("5-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var swords []qualityCalc
	for row := range rows {
		id, err := strconv.Atoi(row[:strings.IndexByte(row, ':')])
		if err != nil {
			panic(err)
		}

		sword := qualityCalc{
			id: id,
			lines: [][3]int{
				{0, 0, 0},
			},
		}

		for numStr := range strings.SplitSeq(row[strings.IndexByte(row, ':')+1:], ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			placed := false
			for i := 0; i < len(sword.lines); i++ {
				if sword.lines[i][1] == 0 {
					sword.lines[i][1] = num
					placed = true
				} else if num < sword.lines[i][1] && sword.lines[i][0] == 0 {
					sword.lines[i][0] = num
					placed = true
					break
				} else if num > sword.lines[i][1] && sword.lines[i][2] == 0 {
					sword.lines[i][2] = num
					placed = true
					break
				}
			}

			if !placed {
				sword.lines = append(sword.lines, [3]int{0, num, 0})
			}
		}

		b := strings.Builder{}
		for i := 0; i < len(sword.lines); i++ {
			b.WriteString(strconv.Itoa(sword.lines[i][1]))
		}

		quality, err := strconv.Atoi(b.String())
		if err != nil {
			panic(err)
		}

		sword.quality = quality
		swords = append(swords, sword)
	}

	slices.SortFunc(swords, func(a, b qualityCalc) int {
		if a.quality == b.quality {
			for i := 0; i < len(a.lines); i++ {
				if a.lines[i][0] != b.lines[i][0] ||
					a.lines[i][1] != b.lines[i][1] ||
					a.lines[i][2] != b.lines[i][2] {

					return (b.lines[i][0] + b.lines[i][1] + b.lines[i][2]) - (a.lines[i][0] + a.lines[i][1] + a.lines[i][2])
				}
			}

			return b.id - a.id
		}

		return b.quality - a.quality
	})

	total := 0
	for i, sword := range swords {
		total += (i + 1) * sword.id
	}

	fmt.Println(total)
}
