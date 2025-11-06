package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

var ruleRegex = regexp.MustCompile("^([A-Z]+):([A-Z,]+)$")

func main() {
	rows, err := util.ReadStrings("11-3", false, "\n")
	if err != nil {
		panic(err)
	}

	rules := make(map[string][]string)

	for row := range rows {
		ruleParts := ruleRegex.FindStringSubmatch(row)

		var consequents []string
		for consequent := range strings.SplitSeq(ruleParts[2], ",") {
			consequents = append(consequents, consequent)
		}

		rules[ruleParts[1]] = consequents
	}

	maxCount := 0
	minCount := math.MaxInt

	for predicate := range rules {
		counts := map[string]int{
			predicate: 1,
		}

		for i := 0; i < 20; i++ {
			nextCounts := make(map[string]int, len(counts))
			for category, count := range counts {
				if consequents, ok := rules[category]; ok {
					for _, consequent := range consequents {
						nextCounts[consequent] += count
					}
				}
			}
			counts = nextCounts
		}

		total := 0
		for _, count := range counts {
			total += count
		}

		if total > maxCount {
			maxCount = total
		}
		if total < minCount {
			minCount = total
		}
	}

	fmt.Println(maxCount - minCount)
}
