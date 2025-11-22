package main

import (
	"bytes"
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

var diagonals = [4][2]int{
	{-1, 1},
	{-1, -1},
	{1, 1},
	{1, -1},
}

type matching struct {
	round int
	gap   int
	score int
}

func main() {
	rows, err := util.ReadStrings("14-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var matches []matching

	var grid [][]byte
	for y := 0; y < 34; y++ {
		grid = append(grid, bytes.Repeat([]byte{'.'}, 34))
	}

	// position the input in the middle of the grid
	var innerGrid [][]byte
	for row := range rows {
		innerGrid = append(innerGrid, []byte(row))
	}

	startY := 17 - len(innerGrid)/2
	startX := 17 - len(innerGrid[0])/2

	prevMatch := 0
	for round := 0; round < 10000; round++ {
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

		matched := true
	checkloop:
		for y := startY; y < startY+len(innerGrid); y++ {
			for x := startX; x < startX+len(innerGrid[0]); x++ {
				if nextGrid[y][x] != innerGrid[y-startY][x-startX] {
					matched = false
					break checkloop
				}
			}
		}

		if matched {
			score := 0
			for y := 0; y < len(nextGrid); y++ {
				for x := 0; x < len(nextGrid[0]); x++ {
					if nextGrid[y][x] == '#' {
						score++
					}
				}
			}

			matches = append(matches, matching{
				round: round + 1,
				gap:   round - prevMatch,
				score: score,
			})

			prevMatch = round
		}

		grid = nextGrid
	}

	total := matches[0].score + matches[1].score

	// we should build a cycle, the cycle being the first time that the same gap appears
	// this all assumes that there isn't some kind of super cycle with the outer pattern
	cycle := []matching{
		matches[1],
	}
	cycleScore := matches[1].score
	cycleGap := matches[1].gap
	for i := 2; i < len(matches); i++ {
		if matches[i].gap == matches[1].gap {
			break
		}

		cycle = append(cycle, matches[i])
		cycleScore += matches[i].score
		cycleGap += matches[i].gap
		total += matches[i].score
	}

	start := cycle[len(cycle)-1].round
	remaining := 1000000000 - start

	m := remaining / cycleGap

	total += cycleScore * m

	rm := remaining % m

	for i := 0; i < len(cycle); i++ {
		if rm > cycle[i].gap {
			total += cycle[i].score
			rm -= cycle[i].gap
		} else {
			break
		}
	}

	fmt.Println(total)
}
