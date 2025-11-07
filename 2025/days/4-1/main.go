package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadInts("4-1", false, "\n")
	if err != nil {
		panic(err)
	}

	ratios := slices.Collect(rows)
	acc := float64(1)
	for i := 1; i < len(ratios); i++ {
		r := float64(ratios[i-1]) / float64(ratios[i])
		acc *= r
	}

	fmt.Println(math.Floor(acc * 2025))
}
