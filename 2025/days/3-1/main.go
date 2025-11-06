package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("3-1", false, "\n")
	if err != nil {
		panic(err)
	}

	sizes := slices.Collect(rows)[0]
	slices.Sort(sizes)

	total := 0
	last := sizes[len(sizes)-1] + 1
	for i := len(sizes) - 1; i >= 0; i-- {
		if sizes[i] < last {
			total += sizes[i]
			last = sizes[i]
		}
	}

	fmt.Println(total)
}
