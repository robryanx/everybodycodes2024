package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadInts("8-1", false, "\n")
	if err != nil {
		panic(err)
	}
	count := slices.Collect(rows)[0]

	blocks := 1
	height := 1
	width := 1

	for blocks < count {
		height++
		width = ((height - 1) * 2) + 1
		blocks += width
	}

	fmt.Println((blocks - count) * width)
}
