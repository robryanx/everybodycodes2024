package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	colsSeq, err := util.ReadInts("11-2", false, "\n")
	if err != nil {
		panic(err)
	}

	cols := slices.Collect(colsSeq)

	// TODO: There is probably an optimisation for the first part as well
	rounds := 0
	for {
		moved := false
		for i := 0; i < len(cols)-1; i++ {
			if cols[i] > cols[i+1] {
				cols[i] -= 1
				cols[i+1] += 1
				moved = true
			}
		}

		if !moved {
			break
		}

		rounds++
	}

	sum := 0
	for i := 0; i < len(cols); i++ {
		sum += cols[i]
	}
	average := sum / len(cols)

	for i := 0; i < len(cols); i++ {
		if cols[i] < average {
			rounds += average - cols[i]
		} else {
			break
		}
	}

	fmt.Println(rounds)
}
