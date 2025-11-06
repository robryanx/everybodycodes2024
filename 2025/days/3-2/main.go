package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("3-2", false, "\n")
	if err != nil {
		panic(err)
	}

	sizes := slices.Collect(rows)[0]
	slices.Sort(sizes)

	count := 0
	last := 0
	total := 0
	for i := 0; i < len(sizes); i++ {
		if sizes[i] > last {
			total += sizes[i]
			last = sizes[i]
			count++

			if count > 19 {
				break
			}
		}
	}

	fmt.Println(total)
}
