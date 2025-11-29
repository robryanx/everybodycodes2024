package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	patternSeq, err := util.ReadInts("16-2", false, ",")
	if err != nil {
		panic(err)
	}

	var result []int
	pattern := slices.Collect(patternSeq)
	rec(pattern, []int{}, 1, &result)

	total := 1
	for _, r := range result {
		total *= r
	}

	fmt.Println(total)
}

func rec(remaingPattern []int, nums []int, next int, result *[]int) {
	checkPattern := slices.Clone(remaingPattern)

	// apply the number
	failed := false
	for i := next - 1; i < len(checkPattern); i += next {
		checkPattern[i]--
		if checkPattern[i] < 0 {
			failed = true
		}
	}

	if !failed {
		finished := true
		for i := 0; i < len(checkPattern); i++ {
			if checkPattern[i] != 0 {
				finished = false
				break
			}
		}

		nums = append(nums, next)
		if finished {
			*result = nums
		} else {
			rec(checkPattern, nums, next+1, result)
		}
	}

	if *result == nil {
		// skip the number
		rec(remaingPattern, nums, next+1, result)
	}
}
