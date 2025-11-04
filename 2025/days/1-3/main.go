package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("1-3", false, "\n")
	if err != nil {
		panic(err)
	}

	lines := slices.Collect(rows)
	names := strings.Split(string(lines[0]), ",")
	for instruction := range strings.SplitSeq(string(lines[2]), ",") {
		pos := 0
		count, err := strconv.Atoi(instruction[1:])
		if err != nil {
			panic(err)
		}
		switch instruction[0] {
		case 'L':
			pos -= count
			pos %= len(names)
			if pos < 0 {
				pos = len(names) + pos
			}
		case 'R':
			pos += count
			pos %= len(names)
		}

		names[0], names[pos] = names[pos], names[0]
	}

	fmt.Println(names[0])
}
