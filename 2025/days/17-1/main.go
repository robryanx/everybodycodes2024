package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("17-1", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	radius := 10
	midY := len(grid) / 2
	midX := len(grid[0]) / 2

	total := 0
	for y := range len(grid) {
		for x := range len(grid[0]) {
			dist := ((midX - x) * (midX - x)) + ((midY - y) * (midY - y))
			if dist == 0 {
				continue
			}
			if dist <= radius*radius {
				total += int(grid[y][x] - '0')
			}
		}
	}

	fmt.Println(total)
}
