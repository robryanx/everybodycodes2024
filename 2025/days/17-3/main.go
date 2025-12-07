package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("17-3", false, "\n")
	if err != nil {
		panic(err)
	}

	var grid [][]byte
	for row := range rows {
		grid = append(grid, []byte(row))
	}

	midY := len(grid) / 2
	midX := len(grid[0]) / 2

	startX := 0
	startY := 0

	for radius := 4; radius < midX-1; radius++ {
		nodeLookup := make(map[int]*node, len(grid)*len(grid[0]))

		for y := range len(grid) {
			for x := range len(grid[0]) {
				if grid[y][x] == '.' {
					continue
				}
				dist := ((midX - x) * (midX - x)) + ((midY - y) * (midY - y))
				if dist == 0 {
					continue
				}
				if dist <= radius*radius {
					grid[y][x] = '.'
				}
				if grid[y][x] == 'S' {
					grid[y][x] = '0'

					startX = x
					startY = y
				}

				n := &node{y, x, int(grid[y][x] - '0'), 0, false}
				nodeLookup[y*1000+x] = n
			}
		}

		start := nodeLookup[startY*1000+startX]
		end := nodeLookup[(midY+radius+1)*1000+midX]
		box := &bounding{0, midX + 2, 4, len(grid[0])}

		firstHalf := pathFind(grid, nodeLookup, start, end, box)

		box = &bounding{0, 0, len(grid), midX - 2}

		for _, node := range nodeLookup {
			node.distance = 0
			node.visited = false
		}

		secondHalf := pathFind(grid, nodeLookup, end, start, box)
		if firstHalf+secondHalf < (radius+1)*30 {
			fmt.Println((firstHalf + secondHalf) * radius)
			break
		}
	}
}

type node struct {
	y        int
	x        int
	val      int
	distance int
	visited  bool
}

type bounding struct {
	y      int
	x      int
	height int
	width  int
}

func pathFind(grid [][]byte, nodeLookup map[int]*node, start, end *node, box *bounding) int {
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

		if curr.x-1 >= 0 && !blocked(curr.y, curr.x-1, box) && grid[curr.y][curr.x-1] != '.' {
			next, ok := nodeLookup[curr.y*1000+curr.x-1]
			if ok && !next.visited {
				dist := curr.distance + next.val
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y-1 >= 0 && !blocked(curr.y-1, curr.x, box) && grid[curr.y-1][curr.x] != '.' {
			next, ok := nodeLookup[(curr.y-1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + next.val
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y+1 < len(grid) && !blocked(curr.y+1, curr.x, box) && grid[curr.y+1][curr.x] != '.' {
			next, ok := nodeLookup[(curr.y+1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + next.val
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.x+1 < len(grid[0]) && !blocked(curr.y, curr.x+1, box) && grid[curr.y][curr.x+1] != '.' {
			next, ok := nodeLookup[curr.y*1000+curr.x+1]
			if ok && !next.visited {
				dist := curr.distance + next.val
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}

		curr.visited = true
	}
}

func blocked(y, x int, box *bounding) bool {
	if box == nil {
		return false
	}

	if x >= box.x &&
		x < box.x+box.width &&
		y >= box.y &&
		y < box.y+box.height {
		return true
	}

	return false
}
