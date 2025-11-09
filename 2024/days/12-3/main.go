package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

func main() {
	rows, err := util.ReadStrings("12-3", false, "\n")
	if err != nil {
		panic(err)
	}

	total := 0
	for row := range rows {
		rowParts := strings.Split(row, " ")
		targetX, err := strconv.Atoi(rowParts[0])
		if err != nil {
			panic(err)
		}
		targetY, err := strconv.Atoi(rowParts[1])
		if err != nil {
			panic(err)
		}

		minTime := targetX/2 - 1
		targetX -= minTime
		targetY -= minTime

		time := 0
		best := math.MaxInt
		maxY := -1
		for targetY >= 0 && targetX > 0 {
			for offsetTime := time; offsetTime >= 0; offsetTime-- {
				powerMin := (time+minTime-offsetTime)/3 - 2
				if powerMin < 1 {
					powerMin = 1
				}

				powerMax := time + minTime - offsetTime
				if powerMax < 1 {
					powerMax = 1
				}

				for power := powerMin; power <= powerMax; power++ {
					checkTime := minTime + time - offsetTime
					checkY := shotY(checkTime, power)
					checkX := minTime + time - offsetTime

					if checkY > targetY && checkX >= targetX {
						break
					}

					if checkY == targetY && checkX == targetX && targetY >= maxY {
						maxY = targetY
						if power < best {
							best = power
						}
					}
					if checkY+1 == targetY && checkX == targetX && targetY >= maxY {
						maxY = targetY
						if power*2 < best {
							best = power * 2
						}
					}
					if checkY+2 == targetY && checkX == targetX && targetY >= maxY {
						maxY = targetY
						if power*3 < best {
							best = power * 3
						}
					}
				}
			}

			if maxY != -1 {
				break
			}

			targetY--
			targetX--
			time++
		}

		total += best
	}

	fmt.Println(total)
}

func shotY(time, power int) int {
	if time <= power {
		return time
	} else if time <= power*2 {
		return power
	} else {
		return power - (time - power*2)
	}
}
