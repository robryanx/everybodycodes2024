package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadIntLists("3-3", false, "\n")
	if err != nil {
		panic(err)
	}

	sizes := slices.Collect(rows)[0]
	
	lists := []map[int]struct{}{
		make(map[int]struct{}),
	}
	for i := 0; i < len(sizes); i++ {
		placed := false
		for j := 0; j < len(lists); j++ {
			if _, ok := lists[j][sizes[i]]; !ok {
				lists[j][sizes[i]] = struct{}{}
				placed = true
				break
			}
		}
		if !placed {
			lists = append(lists, map[int]struct{}{
				sizes[i]: {},
			})
		}
	}

	fmt.Println(len(lists))
}
