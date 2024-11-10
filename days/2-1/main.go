package main

import (
	"fmt"
	"strings"

	"github.com/robryanx/everybodycodes2024/util"
)

func main() {
	strs, err := util.ReadStrings("2-1", false, "\n")
	if err != nil {
		panic(err)
	}

	wordsStr := strs[0][strings.Index(strs[0], ":")+1:]

	count := 0
	for _, word := range strings.Split(wordsStr, ",") {
		count += strings.Count(strs[2], word)
	}

	fmt.Println(count)
}
