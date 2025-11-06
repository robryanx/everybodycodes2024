package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

var ruleRegex = regexp.MustCompile("^([A-Z]):([A-Z,]+)$")

func main() {
	rows, err := util.ReadStrings("11-2", false, "\n")
	if err != nil {
		panic(err)
	}

	rules := make(map[byte][]byte)
	counts := map[byte]int{
		'Z': 1,
	}

	for row := range rows {
		ruleParts := ruleRegex.FindStringSubmatch(row)

		var consequents []byte
		for consequent := range strings.SplitSeq(ruleParts[2], ",") {
			consequents = append(consequents, consequent[0])
		}

		rules[ruleParts[0][0]] = consequents
	}

	for i := 0; i < 10; i++ {
		nextCounts := make(map[byte]int, len(counts))
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

	fmt.Println(total)
}
