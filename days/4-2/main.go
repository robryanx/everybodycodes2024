package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rowsIter, err := util.ReadInts("4-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var nums []int
	lowest := 100000
	for num := range rowsIter {
		if num < lowest {
			lowest = num
		}

		nums = append(nums, num)
	}

	count := 0
	for _, num := range nums {
		count += num - lowest
	}

	fmt.Println(count)
}
