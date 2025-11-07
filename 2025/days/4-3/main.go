package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("4-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var ratios []int
	var altRatios []int

	for row := range rows {
		parts := strings.Split(row, "|")
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		ratios = append(ratios, val)
		if len(parts) > 1 {
			altVal, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			altRatios = append(altRatios, altVal)
		} else {
			altRatios = append(altRatios, 0)
		}
	}

	acc := float64(1)
	for i := 1; i < len(ratios); i++ {
		ratio := ratios[i-1]
		if altRatios[i-1] > 0 {
			ratio = altRatios[i-1]
		}

		r := float64(ratio) / float64(ratios[i])
		acc *= r
	}

	fmt.Printf("%.0f", math.Floor(acc*100))
}
