package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	colsSeq, err := util.ReadInts("11-3", false, "\n")
	if err != nil {
		panic(err)
	}

	cols := slices.Collect(colsSeq)

	sum := 0
	for i := 0; i < len(cols); i++ {
		sum += cols[i]
	}
	average := sum / len(cols)

	total := 0
	for i := 0; i < len(cols); i++ {
		if cols[i] < average {
			total += average - cols[i]
		} else {
			break
		}
	}

	fmt.Println(total)
}
