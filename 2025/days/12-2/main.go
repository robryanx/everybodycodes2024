package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("12-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	val := int(grid[0][0] - '0')
	total += chain(grid, val, 0, 0)

	val = int(grid[len(grid)-1][len(grid[0])-1] - '0')
	total += chain(grid, val, len(grid)-1, len(grid[0])-1)

	fmt.Println(total)
}

var offsets = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func chain(grid [][]byte, val, y, x int) int {
	grid[y][x] = 'x'

	count := 1
	for _, offset := range offsets {
		offsetY := y + offset[0]
		offsetX := x + offset[1]

		if offsetY < 0 ||
			offsetY > len(grid)-1 ||
			offsetX < 0 ||
			offsetX > len(grid[0])-1 ||
			grid[offsetY][offsetX] == 'x' {
			continue
		}

		offsetVal := int(grid[offsetY][offsetX] - '0')
		if offsetVal <= val {
			count += chain(grid, offsetVal, offsetY, offsetX)
		}
	}

	return count
}
