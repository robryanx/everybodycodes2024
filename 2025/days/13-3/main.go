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
	rows, err := util.ReadStrings("13-3", false, "\n")
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

	total++

	rangesOrdered := make([]numRange, len(numRanges))

	frontPointer := 0
	backPointer := len(numRanges) - 1
	for i, n := range numRanges {
		if i%2 == 0 {
			rangesOrdered[frontPointer] = n
			frontPointer++
		} else {
			n.start, n.end = n.end, n.start
			rangesOrdered[backPointer] = n
			backPointer--
		}
	}

	offset := 202520252025 % total
	offsetCheck := 1

	for _, r := range rangesOrdered {
		rangeR := r.end - r.start + 1
		if r.start > r.end {
			rangeR = r.start - r.end + 1
		}

		if offset < offsetCheck+rangeR {
			if r.start > r.end {
				fmt.Println(r.start - (offset - offsetCheck))
			} else {
				fmt.Println(r.start + offset - offsetCheck)
			}

			break
		}

		offsetCheck += rangeR
	}
}
