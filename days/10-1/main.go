package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	rows, err := util.ReadStrings("10-1", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	word := []byte{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == '.' {
				word = append(word, common(grid, y, x))
			}
		}
	}

	fmt.Println(string(word))
}

func common(grid [][]byte, y, x int) byte {
	for i := 0; i < len(grid[y]); i++ {
		if grid[y][i] != '*' && grid[y][i] != '.' && i != x {
			for j := 0; j < len(grid); j++ {
				if grid[j][x] != '*' && grid[j][x] != '.' && j != y {
					if grid[j][x] == grid[y][i] {
						return grid[j][x]
					}
				}
			}
		}
	}

	return '.'
}
