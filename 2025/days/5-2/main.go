package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("5-2", false, "\n")
	if err != nil {
		panic(err)
	}

	weakest := math.MaxInt
	strongest := 0

	for row := range rows {
		lines := [][3]int{
			{0, 0, 0},
		}

		for numStr := range strings.SplitSeq(row[strings.IndexByte(row, ':')+1:], ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			placed := false
			for i := 0; i < len(lines); i++ {
				if lines[i][1] == 0 {
					lines[i][1] = num
					placed = true
				} else if num < lines[i][1] && lines[i][0] == 0 {
					lines[i][0] = num
					placed = true
					break
				} else if num > lines[i][1] && lines[i][2] == 0 {
					lines[i][2] = num
					placed = true
					break
				}
			}

			if !placed {
				lines = append(lines, [3]int{0, num, 0})
			}
		}

		b := strings.Builder{}
		for i := 0; i < len(lines); i++ {
			b.WriteString(strconv.Itoa(lines[i][1]))
		}

		quality, err := strconv.Atoi(b.String())
		if err != nil {
			panic(err)
		}

		if quality > strongest {
			strongest = quality
		}
		if quality < weakest {
			weakest = quality
		}
	}

	fmt.Println(strongest - weakest)
}
