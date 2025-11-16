package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

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

	maxY := len(grid) - 1
	maxX := len(grid[0]) - 1

	lookup := map[string]int{}

	var dragon [2]int
	var sheep [][2]int
	hideout := map[int]struct{}{}
	for y := range grid {
		for x := range grid[0] {
			switch grid[y][x] {
			case 'D':
				dragon = [2]int{y, x}
			case 'S':
				sheep = append(sheep, [2]int{y, x})
			case '#':
				hideout[y*100+x] = struct{}{}
			}
		}
	}

	total := move(TURN_SHEEP, dragon, sheep, maxY, maxX, hideout, lookup)

	fmt.Println(total)
}

type turn int

var (
	TURN_SHEEP  turn = 0
	TURN_DRAGON turn = 1
)

func move(t turn, dragon [2]int, sheep [][2]int, maxY, maxX int, hideout map[int]struct{}, lookup map[string]int) int {
	if t == TURN_SHEEP {
		// do any sheep need to be removed?
		for i := len(sheep) - 1; i >= 0; i-- {
			if sheep[i][0] == dragon[0] &&
				sheep[i][1] == dragon[1] {
				if _, ok := hideout[dragon[0]*100+dragon[1]]; !ok {
					sheep = slices.Delete(sheep, i, i+1)
				}
			}
		}
	}

	terminal, valid := isTerminal(sheep, maxY)

	if terminal {
		if valid {
			return 1
		}

		return 0
	}

	total := 0
	if t == TURN_SHEEP {
		moves := validSheepMoves(dragon, sheep, hideout, maxY)
		if len(moves) == 0 {
			t = TURN_DRAGON
		} else {
			for _, m := range moves {
				nextSheep := slices.Clone(sheep)
				nextSheep[m[0]] = [2]int{m[1], m[2]}

				key := serialise(TURN_DRAGON, dragon, nextSheep)
				count, ok := lookup[key]
				if !ok {
					count = move(TURN_DRAGON, dragon, nextSheep, maxY, maxX, hideout, lookup)
					lookup[key] = count
				}

				total += count
			}
		}
	}

	if t == TURN_DRAGON {
		moves := validDragonMoves(dragon, maxY, maxX)
		for _, m := range moves {
			nextSheep := slices.Clone(sheep)

			key := serialise(TURN_SHEEP, m, nextSheep)
			count, ok := lookup[key]
			if !ok {
				count = move(TURN_SHEEP, m, nextSheep, maxY, maxX, hideout, lookup)
				lookup[key] = count
			}
			total += count
		}
	}

	return total
}

func serialise(t turn, dragon [2]int, sheep [][2]int) string {
	b := strings.Builder{}
	b.WriteString(strconv.Itoa(int(t)))
	b.WriteString("-[")
	b.WriteString(strconv.Itoa(dragon[0]))
	b.WriteString(",")
	b.WriteString(strconv.Itoa(dragon[1]))
	b.WriteString("]-")

	for _, s := range sheep {
		b.WriteString("[")
		b.WriteString(strconv.Itoa(s[0]))
		b.WriteString(",")
		b.WriteString(strconv.Itoa(s[1]))
		b.WriteString("]")
	}

	return b.String()
}

var moves = [][2]int{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
}

func validSheepMoves(curr [2]int, sheep [][2]int, hideout map[int]struct{}, maxY int) [][3]int {
	var valid [][3]int
	for i, s := range sheep {
		if s[0] <= maxY+1 {
			if s[0]+1 == curr[0] && s[1] == curr[1] {
				if _, ok := hideout[curr[0]*100+curr[1]]; !ok {
					continue
				}
			}

			valid = append(valid, [3]int{i, s[0] + 1, s[1]})
		}
	}

	return valid
}

func validDragonMoves(curr [2]int, maxY, maxX int) [][2]int {
	var valid [][2]int
	for _, m := range moves {
		nextY := curr[0] + m[0]
		nextX := curr[1] + m[1]
		if !(nextY < 0 ||
			nextY > maxY ||
			nextX < 0 ||
			nextX > maxX) {
			valid = append(valid, [2]int{nextY, nextX})
		}
	}

	return valid
}

func isTerminal(sheep [][2]int, maxY int) (bool, bool) {
	if len(sheep) == 0 {
		return true, true
	}

	for _, s := range sheep {
		if s[0] > maxY {
			return true, false
		}
	}

	return false, false
}
