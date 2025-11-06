package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("3-1", false, "\n")
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
