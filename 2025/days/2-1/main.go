package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/robryanx/everybodycodes/util"
)

var matricRegex = regexp.MustCompile(`^A=\[(\d+),(\d+)]$`)

type rowMatrix struct {
	x int
	y int
}

func main() {
	rows, err := util.ReadStrings("2-1", false, "\n")
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

	input := rowMatrix{
		x: x,
		y: y,
	}
	result := rowMatrix{}

	for range 3 {
		result = multiply(result, result)
		result = divide(result, rowMatrix{10, 10})
		result = add(result, input)
	}

	fmt.Println(result.String())
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
