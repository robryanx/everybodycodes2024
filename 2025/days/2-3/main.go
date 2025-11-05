package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/robryanx/everybodycodes/util"
)

var matricRegex = regexp.MustCompile(`^A=\[([-\d]+),([-\d]+)]$`)

type rowMatrix struct {
	x int
	y int
}

func main() {
	rows, err := util.ReadStrings("2-3", false, "\n")
	if err != nil {
		panic(err)
	}

	lines := slices.Collect(rows)
	nums := matricRegex.FindStringSubmatch(lines[0])

	x, err := strconv.Atoi(nums[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(nums[2])
	if err != nil {
		panic(err)
	}

	engravedPoint := 0

	input := rowMatrix{x, y}
	for y := range 1001 {
		for x := range 1001 {
			if runCalc(rowMatrix{input.x + x, input.y + y}) {
				engravedPoint++
			}
		}
	}

	fmt.Println(engravedPoint)
}

func runCalc(point rowMatrix) bool {
	result := rowMatrix{}
	for range 100 {
		result = multiply(result, result)
		result = divide(result, rowMatrix{100000, 100000})
		result = add(result, point)

		if result.x < -1000000 ||
			result.x > 1000000 ||
			result.y < -1000000 ||
			result.y > 1000000 {
			return false
		}
	}

	return true
}

func (r rowMatrix) String() string {
	return fmt.Sprintf("[%d,%d]", r.x, r.y)
}

// [X1,Y1] + [X2,Y2] = [X1 + X2, Y1 + Y2]
func add(rowA, rowB rowMatrix) rowMatrix {
	return rowMatrix{
		x: rowA.x + rowB.x,
		y: rowA.y + rowB.y,
	}
}

// [X1,Y1] * [X2,Y2] = [X1 * X2 - Y1 * Y2, X1 * Y2 + Y1 * X2]
func multiply(rowA, rowB rowMatrix) rowMatrix {
	return rowMatrix{
		x: (rowA.x * rowB.x) - (rowA.y * rowB.y),
		y: (rowA.x * rowB.y) + (rowA.y * rowB.x),
	}
}

// [X1,Y1] / [X2,Y2] = [X1 / X2, Y1 / Y2]
func divide(rowA, rowB rowMatrix) rowMatrix {
	return rowMatrix{
		x: rowA.x / rowB.x,
		y: rowA.y / rowB.y,
	}
}
