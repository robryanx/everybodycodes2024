package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadInts("8-2", false, "\n")
	if err != nil {
		panic(err)
	}
	priests := slices.Collect(rows)[0]
	blockCount := 20240000

	blocks := 1
	step := 1
	width := 1
	heightAdd := 1

	for blocks < blockCount {
		step++
		heightAdd = (heightAdd * priests) % 1111
		width = ((step - 1) * 2) + 1
		blocks += width * heightAdd
	}

	fmt.Println((blocks - blockCount) * width)
}
