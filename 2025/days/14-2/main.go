package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

var diagonals = [4][2]int{
	{-1, 1},
	{-1, -1},
	{1, 1},
	{1, -1},
}

func main() {
	rows, err := util.ReadStrings("14-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	total := 0
	for round := 0; round < 2025; round++ {
		nextGrid := util.CopyGrid(grid, true)
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				active := 0
				for _, diagonal := range diagonals {
					checkY := y + diagonal[0]
					checkX := x + diagonal[1]

					if checkY < 0 ||
						checkY > len(grid)-1 ||
						checkX < 0 ||
						checkX > len(grid[0])-1 {
						continue
					}

					if grid[checkY][checkX] == '#' {
						active++
					}
				}

				if active%2 == 0 {
					if grid[y][x] == '#' {
						nextGrid[y][x] = '.'
					} else {
						nextGrid[y][x] = '#'
					}
				}
			}
		}

		for y := 0; y < len(nextGrid); y++ {
			for x := 0; x < len(nextGrid[0]); x++ {
				if nextGrid[y][x] == '#' {
					total++
				}
			}
		}

		grid = nextGrid
	}

	fmt.Println(total)
}
