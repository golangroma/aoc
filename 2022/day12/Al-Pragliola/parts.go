package main

import (
	"strconv"
)

func PartOne(input []string) string {
	w := ParseWorld(input)

	p, _, found := Path(w.From(), w.To())

	if !found {
		return "no path found"
	}

	return strconv.Itoa(len(p) - 1)
}

func PartTwo(input []string) string {
	result := 1<<63 - 1

	w := ParseWorld(input)

	for y, row := range input {
		for x, raw := range row {
			if raw == 'a' || raw == 'S' {
				p, _, found := Path(w.Tile(x, y), w.To())

				if !found {
					continue
				}

				if len(p)-1 < result {
					result = len(p) - 1
				}
			}
		}
	}

	return strconv.Itoa(result)
}
