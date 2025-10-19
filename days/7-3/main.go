package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes2024/util"
)

var trackStr = `S+= +=-== +=++=     =+=+=--=    =-= ++=     +=-  =+=++=-+==+ =++=-=-=--
- + +   + =   =     =      =   == = - -     - =  =         =-=        -
= + + +-- =-= ==-==-= --++ +  == == = +     - =  =    ==++=    =++=-=++
+ + + =     +         =  + + == == ++ =     = =  ==   =   = =++=
= = + + +== +==     =++ == =+=  =  +  +==-=++ =   =++ --= + =
+ ==- = + =   = =+= =   =       ++--          +     =   = = =--= ==++==
=     ==- ==+-- = = = ++= +=--      ==+ ==--= +--+=-= ==- ==   =+=    =
-               = = = =   +  +  ==+ = = +   =        ++    =          -
-               = + + =   +  -  = + = = +   =        +     =          -
--==++++==+=+++-= =-= =-+-=  =+-= =-= =--   +=++=+++==     -=+=++==+++-`

func main() {
	rows, err := util.ReadStrings("7-3", false, "\n")
	if err != nil {
		panic(err)
	}
	row := slices.Collect(rows)[0]
	actions := parseActions(row[1:])
	track := parseTrack(trackStr)
	baseLine := scoreActions(actions, track)

	var actionList [][]action
	remainingPlusActions := 5
	remainingMinusActions := 3
	remainingEqualsActions := 3

	actionBuild(remainingPlusActions, remainingMinusActions, remainingEqualsActions, []action{}, &actionList)

	winningPlans := 0
	for _, checkActions := range actionList {
		if scoreActions(checkActions, track) > baseLine {
			winningPlans++
		}
	}

	fmt.Println(winningPlans)
}

func actionBuild(remainingPlusActions, remainingMinusActions, remainingEqualsActions int, current []action, actionList *[][]action) {
	if remainingPlusActions == 0 && remainingMinusActions == 0 && remainingEqualsActions == 0 {
		*actionList = append(*actionList, current)
		return
	}

	if remainingPlusActions > 0 {
		next := slices.Clone(current)
		next = append(next, actionPlus)
		actionBuild(remainingPlusActions-1, remainingMinusActions, remainingEqualsActions, next, actionList)
	}
	if remainingMinusActions > 0 {
		next := slices.Clone(current)
		next = append(next, actionMinus)
		actionBuild(remainingPlusActions, remainingMinusActions-1, remainingEqualsActions, next, actionList)
	}
	if remainingEqualsActions > 0 {
		next := slices.Clone(current)
		next = append(next, actionEquals)
		actionBuild(remainingPlusActions, remainingMinusActions, remainingEqualsActions-1, next, actionList)
	}
}

type action int8

const (
	actionEmpty  action = 0
	actionPlus   action = 1
	actionMinus  action = 2
	actionEquals action = 3
	actionStart  action = 4
)

func scoreActions(actions []action, track []action) int {
	score := 0
	count := 0
	current := 10
	for range 11 {
		for _, t := range track {
			switch t {
			case actionPlus:
				current++
			case actionMinus:
				current--
			default:
				switch actions[count%len(actions)] {
				case actionPlus:
					current++
				case actionMinus:
					current--
				}
			}
			count++
			score += current
		}
	}

	return score
}

func parseActions(plan string) []action {
	var actions []action

	for _, p := range plan {
		switch p {
		case '=':
			actions = append(actions, actionEquals)
		case '+':
			actions = append(actions, actionPlus)
		case '-':
			actions = append(actions, actionMinus)
		}
	}

	return actions
}

func printTrack(track []action) {
	for _, a := range track {
		switch a {
		case actionPlus:
			fmt.Print("+")
		case actionMinus:
			fmt.Print("-")
		case actionEquals:
			fmt.Print("=")
		case actionStart:
			fmt.Print("S")
		}
	}
}

func parseTrack(trackStr string) []action {
	var track []action
	var grid [][]action
	rows := strings.Split(trackStr, "\n")
	for _, row := range rows {
		buildRow := make([]action, len(rows[0]))
		for i, ch := range row {
			switch ch {
			case '=':
				buildRow[i] = actionEquals
			case '+':
				buildRow[i] = actionPlus
			case '-':
				buildRow[i] = actionMinus
			case 'S':
				buildRow[i] = actionStart
			}
		}
		grid = append(grid, buildRow)
	}

	prevY := 0
	prevX := 0
	currY := 0
	currX := 1
	for grid[currY][currX] != actionStart {
		track = append(track, grid[currY][currX])

		if currY-1 >= 0 && currY-1 != prevY && grid[currY-1][currX] != actionEmpty {
			prevY = currY
			prevX = currX
			currY -= 1
		} else if currY+1 < len(grid) && currY+1 != prevY && grid[currY+1][currX] != actionEmpty {
			prevY = currY
			prevX = currX
			currY += 1
		} else if currX-1 >= 0 && currX-1 != prevX && grid[currY][currX-1] != actionEmpty {
			prevX = currX
			prevY = currY
			currX -= 1
		} else if currX+1 < len(grid[0]) && currX+1 != prevX && grid[currY][currX+1] != actionEmpty {
			prevX = currX
			prevY = currY
			currX += 1
		}
	}

	track = append(track, grid[0][0])

	return track
}
