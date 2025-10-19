package main

import (
	"fmt"
	"slices"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadInts("8-3", false, "\n")
	if err != nil {
		panic(err)
	}
	priests := slices.Collect(rows)[0]
	acolytes := 10
	blockCount := 202400000

	blocks := 1
	step := 1
	width := 1
	heightAdd := 1

	var heightAdditions []int

	for blocks < blockCount {
		step++
		heightAdd = ((heightAdd * priests) % acolytes) + acolytes
		width = ((step - 1) * 2) + 1
		blocks += width * heightAdd
		heightAdditions = append(heightAdditions, heightAdd)
	}

	remove := 0
	height := 0
	for i := 1; i < (width / 2); i++ {
		height = 0
		for j := len(heightAdditions) - 1; j > len(heightAdditions)-2-i; j-- {
			height += heightAdditions[j]
		}
		remove += ((priests * width * height) % acolytes) * 2
	}

	remove += ((priests * width * (height + 1)) % acolytes)

	blocks -= remove

	fmt.Println(blocks - blockCount)
}
