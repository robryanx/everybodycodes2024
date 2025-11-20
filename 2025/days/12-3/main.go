package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("12-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var baseGrid [][]byte
	for row := range rows {
		baseGrid = append(baseGrid, []byte(row))
	}

	total := 0
	var bestGrid [][]byte

	// TODO: Optimise this
	for z := range 3 {
		best := 0
		useGrid := baseGrid
		if z != 0 {
			useGrid = bestGrid
		}

		for y := 0; y < len(useGrid); y++ {
			for x := 0; x < len(useGrid[0]); x++ {
				if useGrid[y][x] == 'x' {
					continue
				}

				val := int(useGrid[y][x] - '0')
				grid := util.CopyGrid(useGrid, true)
				check := chain(grid, val, y, x)
				if check > best {
					best = check
					bestGrid = grid
				}
			}
		}

		total += best
	}

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
