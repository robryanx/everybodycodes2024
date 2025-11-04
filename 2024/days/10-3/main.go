package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("10-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	offsetY := 0
	offsetX := 0

	total := 0

	gridsX := (len(grid[0]) + 1) / 6
	gridsY := (len(grid) + 1) / 6

	for x := range 2 {
		for gridY := 0; gridY < gridsY; gridY++ {
			for gridX := 0; gridX < gridsX; gridX++ {
				offsetY = gridY * 6
				offsetX = gridX * 6

				for y := offsetY; y < offsetY+8; y++ {
					for x := offsetX; x < offsetX+8; x++ {
						if grid[y][x] == '.' {
							grid[y][x] = common(grid, offsetY, offsetX, y, x)
						}
					}
				}

				if x == 0 {
					for y := offsetY; y < offsetY+8; y++ {
						for x := offsetX; x < offsetX+8; x++ {
							if grid[y][x] == '.' {
								unmatched(grid, offsetY, offsetX, y, x)
							}
						}
					}
				}

				if x == 1 {
					count := 1
					blockTotal := 0
					skip := false
				skipLabel:
					for y := offsetY + 2; y < offsetY+6; y++ {
						for x := offsetX + 2; x < offsetX+6; x++ {
							if grid[y][x] == '.' {
								skip = true
								break skipLabel
							}

							blockTotal += int(grid[y][x]-'A'+1) * count
							count++
						}
					}

					if !skip {
						total += blockTotal
					}
				}

			}
		}
	}

	fmt.Println(total)
}

func unmatched(grid [][]byte, offsetY, offsetX, y, x int) {
	// which col contains the question mark?
	questionY := -1
	questionX := -1

	valCount := make(map[byte]int, 24)
	for i := offsetX; i < offsetX+8; i++ {
		if grid[y][i] == '?' {
			questionY = y
			questionX = i
		} else if grid[y][i] != '.' {
			valCount[grid[y][i]]++
		}
	}

	for j := offsetY; j < offsetY+8; j++ {
		if grid[j][x] == '?' {
			questionY = j
			questionX = x
		} else if grid[j][x] != '.' {
			valCount[grid[j][x]]++
		}
	}

	if questionY != -1 {
		matches := 0
		matchCh := byte('0')
		for ch, count := range valCount {
			if count == 1 {
				matches++
				if matches > 1 {
					break
				}
				matchCh = ch
			}
		}
		if matches == 1 {
			grid[y][x] = matchCh
			grid[questionY][questionX] = matchCh
		}
	}
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

			if grid[j][x] == grid[y][i] {
				return grid[j][x]
			}
		}
	}

	return '.'
}
