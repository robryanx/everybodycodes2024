package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("5-1", false, "\n")
	if err != nil {
		panic(err)
	}

	lines := [][3]int{
		{0, 0, 0},
	}
	row := slices.Collect(rows)[0]
	for numStr := range strings.SplitSeq(row[strings.IndexByte(row, ':')+1:], ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		placed := false
		for i := 0; i < len(lines); i++ {
			if lines[i][1] == 0 {
				lines[i][1] = num
				placed = true
			} else if num < lines[i][1] && lines[i][0] == 0 {
				lines[i][0] = num
				placed = true
				break
			} else if num > lines[i][1] && lines[i][2] == 0 {
				lines[i][2] = num
				placed = true
				break
			}
		}

		if !placed {
			lines = append(lines, [3]int{0, num, 0})
		}
	}

	b := strings.Builder{}
	for i := 0; i < len(lines); i++ {
		b.WriteString(strconv.Itoa(lines[i][1]))
	}

	fmt.Println(b.String())
}
