package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadInts("13-1", false, "\n")
	if err != nil {
		panic(err)
	}

	nums := slices.Collect(rows)

	ordered := make([]int, len(nums)+1)
	ordered[0] = 1

	frontPointer := 1
	backPointer := len(nums)
	for i := range len(nums) {
		if i%2 == 0 {
			ordered[frontPointer] = nums[i]
			frontPointer++
		} else {
			ordered[backPointer] = nums[i]
			backPointer--
		}
	}

	fmt.Println(ordered[2025%len(ordered)])
}
