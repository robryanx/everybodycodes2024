package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	strs, err := util.ReadBytes("1-2", false)
	if err != nil {
		panic(err)
	}

	sum := 0
	for pairs := range slices.Chunk(strs, 2) {
		isPair := true
		for _, char := range pairs {
			switch char {
			case 'B':
				sum++
			case 'C':
				sum += 3
			case 'D':
				sum += 5
			case 'x':
				isPair = false
			}
		}

		if isPair {
			sum += 2
		}
	}

	fmt.Println(sum)
}
