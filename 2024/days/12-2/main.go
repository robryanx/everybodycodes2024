package main

import (
	"fmt"

	"github.com/robryanx/everybodycodes/util"
)

type pos struct {
	segment int
	y       int
	x       int
}

func main() {
	rows, err := util.ReadStrings("12-2", false, "\n")
	if err != nil {
		panic(err)
	}

	var catapults []pos
	var targets []pos

	y := 0
	for row := range rows {
		for x, val := range []byte(row) {
			if val >= 'A' && val <= 'C' {
				catapults = append(catapults, pos{int(val - '@'), y, x})
			} else if val == 'T' || val == 'H' {
				targets = append(targets, pos{int(val - '@'), y, x})
			}
		}

		y++
	}

	total := 0
	for _, target := range targets {
		for _, catapult := range catapults {
			if power := canHit(catapult, target); power != 0 {
				if target.segment == 8 {
					total += catapult.segment * 2 * power
				} else {
					total += catapult.segment * power
				}

				break
			}
		}
	}

	fmt.Println(total)
}

func canHit(catapult, target pos) int {
	xDiff := target.x - catapult.x
	yDiff := target.y - catapult.y

	power := 1
	for {
		path := power*3 + yDiff

		if path == xDiff {
			return power
		} else if path > xDiff {
			return 0
		}
		power++
	}
}
