package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("3-3", false, "\n")
	if err != nil {
		panic(err)
	}

	list := slices.Collect(rows)[0]
	sizes := make([]int, 0, 124)
	for size := range strings.SplitSeq(list, ",") {
		sizeParsed, err := strconv.Atoi(size)
		if err != nil {
			panic(err)
		}
		sizes = append(sizes, sizeParsed)
	}

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
