package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

type node struct {
	y        int
	x        int
	distance int
	visited  bool
}

type facing int32

const (
	facingNorth facing = iota
	facingEast
	facingSouth
	facingWest
)

type instruction struct {
	steps     int
	direction facing
}

func main() {
	rows, err := util.ReadStrings("15-2", false, "\n")
	if err != nil {
		panic(err)
	}

	currentY := 0
	currentX := 0
	minY := 0
	maxY := 0
	minX := 0
	maxX := 0

	var instructions []instruction

	currentFacing := facingNorth
	for instructionStr := range strings.SplitSeq(slices.Collect(rows)[0], ",") {
		steps, err := strconv.Atoi(instructionStr[1:])
		if err != nil {
			panic(err)
		}

		currentFacing = nextFacing(currentFacing, instructionStr[:1])

		instructions = append(instructions, instruction{
			steps:     steps,
			direction: currentFacing,
		})

		switch currentFacing {
		case facingNorth:
			currentY -= steps
			if currentY < minY {
				minY = currentY
			}
		case facingEast:
			currentX += steps
			if currentX > maxX {
				maxX = currentX
			}
		case facingSouth:
			currentY += steps
			if currentY > maxY {
				maxY = currentY
			}
		case facingWest:
			currentX -= steps
			if currentX < minX {
				minX = currentX
			}
		}
	}

	yLen := (maxY - minY) + 3
	startingY := -minY + 1

	xLen := (maxX - minX) + 3
	startingX := -minX + 1

	var grid [][]byte
	for y := 0; y < yLen; y++ {
		grid = append(grid, bytes.Repeat([]byte("."), xLen))
	}

	currentY = startingY
	currentX = startingX

	for _, inst := range instructions {
		switch inst.direction {
		case facingNorth:
			for y := currentY; y >= currentY-inst.steps; y-- {
				grid[y][currentX] = '#'
			}
			currentY -= inst.steps
		case facingEast:
			for x := currentX; x < currentX+inst.steps; x++ {
				grid[currentY][x] = '#'
			}
			currentX += inst.steps
		case facingSouth:
			for y := currentY; y < currentY+inst.steps; y++ {
				grid[y][currentX] = '#'
			}
			currentY += inst.steps
		case facingWest:
			for x := currentX; x >= currentX-inst.steps; x-- {
				grid[currentY][x] = '#'
			}
			currentX -= inst.steps
		}
	}

	nodeLookup := make(map[int]*node, len(grid)*len(grid[0]))

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] != '#' {
				n := &node{y, x, 0, false}
				nodeLookup[y*1000+x] = n
			}
		}
	}

	start := &node{startingY, startingX, 0, false}
	nodeLookup[startingY*1000+startingX] = start
	grid[startingY][startingX] = 'S'
	end := &node{currentY, currentX, 0, false}
	nodeLookup[currentY*1000+currentX] = end
	grid[currentY][currentX] = 'E'

	fmt.Println(pathFind(grid, nodeLookup, start, end))
}

func nextFacing(currentFacing facing, direction string) facing {
	switch currentFacing {
	case facingNorth:
		if direction == "L" {
			return facingWest
		}
		return facingEast
	case facingEast:
		if direction == "L" {
			return facingNorth
		}
		return facingSouth
	case facingSouth:
		if direction == "L" {
			return facingEast
		}
		return facingWest
	case facingWest:
		if direction == "L" {
			return facingSouth
		}
		return facingNorth
	}
	return -1
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
				dist := curr.distance + 1
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y-1 >= 0 {
			next, ok := nodeLookup[(curr.y-1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + 1
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.y+1 < len(grid) {
			next, ok := nodeLookup[(curr.y+1)*1000+curr.x]
			if ok && !next.visited {
				dist := curr.distance + 1
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}
		if curr.x+1 < len(grid[0]) {
			next, ok := nodeLookup[curr.y*1000+curr.x+1]
			if ok && !next.visited {
				dist := curr.distance + 1
				if dist < next.distance || next.distance == 0 {
					next.distance = dist
					pq.Push(next)
				}
			}
		}

		curr.visited = true
	}
}
