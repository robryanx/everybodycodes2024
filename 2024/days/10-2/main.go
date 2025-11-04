package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("10-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	gridsX := (len(grid[0]) + 1) / 9
	gridsY := (len(grid) + 1) / 9

	total := 0
	for gridY := 0; gridY < gridsY; gridY++ {
		for gridX := 0; gridX < gridsX; gridX++ {
			offsetY := gridY * 9
			offsetX := gridX * 9

			wordTotal := 0
			count := 1
			for y := offsetY; y < offsetY+8; y++ {
				for x := offsetX; x < offsetX+8; x++ {
					if grid[y][x] == '.' {
						wordTotal += int(common(grid, offsetY, offsetX, y, x)-'A'+1) * count
						count++
					}
				}
			}

			total += wordTotal
		}
	}

	fmt.Println(total)
}

func common(grid [][]byte, offsetY, offsetX, y, x int) byte {
	for i := offsetX; i < offsetX+8; i++ {
		if i >= offsetX+2 && i < offsetX+6 {
			continue
		}

		for j := offsetY; j < offsetY+8; j++ {
			if j >= offsetY+2 && j < offsetY+6 {
				continue
			}

			if j != y && grid[j][x] == grid[y][i] {
				return grid[j][x]
			}
		}
	}

	return '.'
}
