package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("8-1", false, "\n")
	if err != nil {
		panic(err)
	}

	row := slices.Collect(rows)[0]

	total := 0
	nailCount := 32
	testCount := nailCount / 2
	for i := 0; i < len(row)-1; i++ {
		if abs(row[i]-row[i+1]) == testCount {
			total++
		}
	}

	fmt.Println(total)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
