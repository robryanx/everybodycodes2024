package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("13-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	var start *node
	var end *node

	nodeLookup := make(map[int]*node, len(grid)*len(grid[0]))

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '#' {
				n := &node{y, x, height(grid[y][x]), 0, false}
				switch grid[y][x] {
				case 'S':
					start = n
				case 'E':
					end = n
				}
				nodeLookup[y*1000+x] = n
			}
		}
	}

	fmt.Println(pathFind(grid, nodeLookup, start, end))
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

		if curr.x-1 >= 0 {
			next, ok := nodeLookup[curr.y*1000+curr.x-1]
			if ok && !next.visited {
				dist := curr.distance + distance(curr.height, next.height)
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y-1 >= 0 {
			next, ok := nodeLookup[(curr.y-1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + distance(curr.height, next.height)
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y+1 < len(grid) {
			next, ok := nodeLookup[(curr.y+1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + distance(curr.height, next.height)
				if dist < next.distance || next.distance == 0 {
					next.distance = curr.distance + distance(curr.height, next.height)
					pq.Push(next)
				}
			}
		}
		if curr.x+1 < len(grid[0]) {
			next, ok := nodeLookup[curr.y*1000+curr.x+1]
			if ok && !next.visited {
				dist := curr.distance + distance(curr.height, next.height)
				if dist < next.distance || next.distance == 0 {
					next.distance = curr.distance + distance(curr.height, next.height)
					pq.Push(next)
				}
			}
		}

		curr.visited = true
	}
}
