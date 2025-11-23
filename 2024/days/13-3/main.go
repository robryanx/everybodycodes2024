package main

import (
	"fmt"
	"math"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("13-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	var end *node

	nodeLookup := make(map[int]*node, len(grid)*len(grid[0]))

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '#' {
				n := &node{y, x, height(grid[y][x]), 0, false}
				switch grid[y][x] {
				case 'E':
					end = n
				}
				nodeLookup[y*1000+x] = n
			}
		}
	}

	_ = pathFind(grid, nodeLookup, end, nil)

	best := math.MaxInt
	for _, y := range []int{0, len(grid) - 1} {
		for x := 1; x < len(grid[0])-1; x++ {
			if nodeLookup[y*1000+x].distance < best {
				best = nodeLookup[y*1000+x].distance
			}
		}
	}

	for _, x := range []int{0, len(grid[0]) - 1} {
		for y := 1; y < len(grid)-1; y++ {
			if nodeLookup[y*1000+x].distance < best {
				best = nodeLookup[y*1000+x].distance
			}
		}
	}

	fmt.Println(best)
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func distance(heightA, heightB int) int {
	dist := abs(heightA - heightB)
	if dist > 5 {
		dist = 10 - dist
	}
	return dist + 1
}

func height(ch byte) int {
	if ch >= '0' && ch <= '9' {
		return int(ch - '0')
	}
	return 0
}

func distanceMap(grid [][]byte, nodeLookup map[int]*node) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			next, ok := nodeLookup[y*1000+x]
			if !ok {
				fmt.Print("##")
			} else {
				fmt.Printf("%02d", next.distance)
			}
		}
		fmt.Println("")
	}
}

type node struct {
	y        int
	x        int
	height   int
	distance int
	visited  bool
}

var offsets = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func pathFind(grid [][]byte, nodeLookup map[int]*node, start, end *node) int {
	pq := util.NewPriorityQueue([]*node{start}, func(a, b *node) bool {
		return a.distance < b.distance
	})

	for {
		curr, ok := pq.Pop()
		if !ok {
			return -1
		}
		if curr == end {
			return curr.distance
		}

		for _, offset := range offsets {
			if curr.x+offset[1] >= 0 &&
				curr.x+offset[1] < len(grid[0]) &&
				curr.y+offset[0] >= 0 &&
				curr.y+offset[0] < len(grid) {

				next, ok := nodeLookup[(curr.y+offset[0])*1000+curr.x+offset[1]]
				if ok && !next.visited {
					dist := curr.distance + distance(curr.height, next.height)
					if dist < next.distance || next.distance == 0 {
						next.distance = dist
						pq.Push(next)
					}
				}
			}
		}

		curr.visited = true
	}
}
