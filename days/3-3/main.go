package main

import (
	"fmt"
	"iter"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rowsIter, err := util.ReadStrings("3-3", false, "\n")
	if err != nil {
		panic(err)
	}

	grid := gridFromRowsIter(rowsIter)

	count := countDug(grid)

	for {
		nextCount := 0
		nextGrid := util.CopyGrid(grid, false)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[0]); x++ {
				if grid[y][x] == '.' {
					nextGrid[y][x] = '.'
				} else {
					dirMatches := 0
					util.AdjacentMatch(grid, y, x, true, func(char byte, y, x int) bool {
						if char != '#' {
							return false
						}

						if char == '#' {
							dirMatches++
						}

						return dirMatches == 8
					})

					if dirMatches == 8 {
						nextGrid[y][x] = '#'
						nextCount++
					} else {
						nextGrid[y][x] = '.'
					}
				}
			}
		}

		count += nextCount
		if nextCount < 9 {
			break
		}

		grid = nextGrid
	}

	fmt.Println(count)
}

func gridFromRowsIter(rowsIter iter.Seq[string]) [][]byte {
	var grid [][]byte
	for row := range rowsIter {
		grid = append(grid, []byte(row))
	}

	return grid
}

func countDug(grid [][]byte) int {
	var count int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				count++
			}
		}
	}

	return count
}
