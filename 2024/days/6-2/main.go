package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

type node struct {
	val     string
	parents []string
}

func main() {
	rows, err := util.ReadStrings("6-2", false, "\n")
	if err != nil {
		panic(err)
	}

	treeLookup := map[string][]string{}
	for row := range rows {
		node, children, _ := strings.Cut(row, ":")
		treeLookup[node] = strings.Split(children, ",")
	}

	lastDepth := 0
	fruitAtDepth := 0
	lastFruit := ""

	nodeList := []node{
		{
			val: "RR",
		},
	}

	for i := 0; ; i++ {
		if len(nodeList) == 0 {
			break
		}

		n := nodeList[0]
		if len(n.parents) > lastDepth {
			if fruitAtDepth == 1 {
				fmt.Println(lastFruit)
				break
			}

			lastDepth = len(n.parents)
			fruitAtDepth = 0
			lastFruit = ""
		}
		nodeList = nodeList[1:]

		for _, nextNode := range treeLookup[n.val] {
			if nextNode == "@" {
				fruitAtDepth++

				var sb strings.Builder
				sb.Grow(len(n.parents) + 2)
				for _, parentNode := range n.parents {
					sb.WriteByte(parentNode[0])
				}
				sb.WriteByte(n.val[0])
				sb.WriteByte('@')
				lastFruit = sb.String()
			}

			nodeList = append(nodeList, node{
				val:     nextNode,
				parents: append(slices.Clone(n.parents), n.val),
			})
		}
	}
}
