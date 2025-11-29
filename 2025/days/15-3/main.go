package main

import (
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
	startY int
	startX int
	endY   int
	endX   int
}

func main() {
	rows, err := util.ReadStrings("15-3", false, "\n")
	if err != nil {
		panic(err)
	}

	currentY := 0
	currentX := 0

	var instructions []instruction

	currentFacing := facingNorth
	for instructionStr := range strings.SplitSeq(slices.Collect(rows)[0], ",") {
		steps, err := strconv.Atoi(instructionStr[1:])
		if err != nil {
			panic(err)
		}

		inst := instruction{
			startY: currentY,
			startX: currentX,
		}

		currentFacing = nextFacing(currentFacing, instructionStr[:1])

		switch currentFacing {
		case facingNorth:
			currentY -= steps
		case facingEast:
			currentX += steps
		case facingSouth:
			currentY += steps
		case facingWest:
			currentX -= steps
		}

		inst.endX = currentX
		inst.endY = currentY

		instructions = append(instructions, inst)
	}

	fmt.Println(instructions)

	// TODO: How to path find when the grid is too big to pathfind
	// We could would how how to work out if our path is going to intersect with any line and the planned path just be the most nieve where possible
	// Division gives us an approx value - we could work out how to map that path with the real values
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
