package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	patternSeq, err := util.ReadInts("16-3", false, ",")
	if err != nil {
		panic(err)
	}

	var result []int
	pattern := slices.Collect(patternSeq)
	rec(pattern, []int{}, 1, &result)

	totalBlocks := 202520252025000
	guess := totalBlocks / 2
	lastGuess := 0
	for {
		totalCh := 0
		for _, r := range result {
			totalCh += guess / r
		}

		if totalCh > totalBlocks {
			guess -= (totalCh - totalBlocks) / 2
		} else {
			guess += (totalBlocks - totalCh) / 2
		}

		if guess == lastGuess {
			if totalCh > totalBlocks {
				guess--
			}

			break
		}
		lastGuess = guess
	}

	fmt.Println(guess)
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
