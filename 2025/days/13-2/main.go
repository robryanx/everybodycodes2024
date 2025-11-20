package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

type numRange struct {
	start int
	end   int
}

func main() {
	rows, err := util.ReadStrings("13-2", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	var numRanges []numRange
	for row := range rows {
		parts := strings.Split(row, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		total += end - start + 1
		numRanges = append(numRanges, numRange{start, end})
	}

	ordered := make([]int, total+1)
	ordered[0] = 1

	frontPointer := 1
	backPointer := len(ordered) - 1

	for i, n := range numRanges {
		if i%2 == 0 {
			for j := n.start; j <= n.end; j++ {
				ordered[frontPointer] = j
				frontPointer++
			}
		} else {
			for j := n.start; j <= n.end; j++ {
				ordered[backPointer] = j
				backPointer--
			}
		}
	}

	fmt.Println(ordered[20252025%len(ordered)])
}
