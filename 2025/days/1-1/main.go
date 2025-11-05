package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("1-1", false, "\n")
	if err != nil {
		panic(err)
	}

	lines := slices.Collect(rows)
	names := strings.Split(lines[0], ",")

	pos := 0
	for instruction := range strings.SplitSeq(lines[2], ",") {
		count := int(instruction[1] - '0')
		switch instruction[0] {
		case 'L':
			pos -= count
			if pos < 0 {
				pos = 0
			}
		case 'R':
			pos += count
			if pos > len(names)-1 {
				pos = len(names) - 1
			}
		}
	}

	fmt.Println(names[pos])
}
