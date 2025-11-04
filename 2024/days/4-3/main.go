package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rowsIter, err := util.ReadInts("4-3", false, "\n")
	if err != nil {
		panic(err)
	}

	nums := slices.Collect(rowsIter)

	// take the median value
	slices.Sort(nums)
	targetHeight := nums[250]

	count := 0
	for _, num := range nums {
		count += abs(num - targetHeight)
	}

	fmt.Println(count)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}

	return x
}
