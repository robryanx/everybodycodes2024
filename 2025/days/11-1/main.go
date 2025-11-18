package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	colsSeq, err := util.ReadInts("11-1", false, "\n")
	if err != nil {
		panic(err)
	}

	cols := slices.Collect(colsSeq)

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

	for i := rounds; i < 10; i++ {
		for i := 0; i < len(cols)-1; i++ {
			if cols[i] < cols[i+1] {
				cols[i] += 1
				cols[i+1] -= 1
			}
		}
	}

	total := 0
	for i := 0; i < len(cols); i++ {
		total += (i + 1) * cols[i]
	}

	fmt.Println(total)
}
