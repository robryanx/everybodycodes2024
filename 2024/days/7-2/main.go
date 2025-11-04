package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/robryanx/everybodycodes/util"
)

type device struct {
	id    byte
	power int
}

var testTrack = `S+===
-   +
=+=-+`

var track = `S-=++=-==++=++=-=+=-=+=+=--=-=++=-==++=-+=-=+=-=+=+=++=-+==++=++=-=-=--
-                                                                     -
=                                                                     =
+                                                                     +
=                                                                     +
+                                                                     =
=                                                                     =
-                                                                     -
--==++++==+=+++-=+=-=+=-+-=+-=+-=+=-=+=--=+++=++=+++==++==--=+=++==+++-`

func parseTrack(track string) string {
	parts := strings.Split(track, "\n")

	buildTrack := parts[0]
	for y := 1; y < len(parts)-1; y++ {
		buildTrack += string(parts[y][len(parts[y])-1])
	}
	st := []byte(parts[len(parts)-1])
	slices.Reverse(st)
	buildTrack += string(st)
	for y := len(parts) - 2; y > 0; y-- {
		buildTrack += string(parts[y][0])
	}

	return buildTrack[1:] + "S"
}

func main() {
	rows, err := util.ReadStrings("7-2", false, "\n")
	if err != nil {
		panic(err)
	}

	tr := parseTrack(track)

	results := make([]device, 0, 100)
	for row := range rows {
		total := 0
		current := 10
		id := row[0]

		segments := (len(row) - 1) / 2

		for i := 0; i < 10; i++ {
			for j, pos := range tr {
				if pos == '+' {
					current++
				} else if pos == '-' {
					current--
				} else {
					segment := row[2+((i*len(tr)+j)%segments)*2]
					if segment == '+' {
						current++
					} else if segment == '-' {
						current--
					}
				}

				total += current
			}
		}

		results = append(results, device{
			id:    id,
			power: total,
		})
	}

	slices.SortFunc(results, func(a, b device) int {
		return cmp.Compare(b.power, a.power)
	})

	for _, r := range results {
		fmt.Print(string(r.id))
	}
	fmt.Println()
}
