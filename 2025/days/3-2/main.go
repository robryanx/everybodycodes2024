package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("3-2", false, "\n")
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
