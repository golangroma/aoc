package main

import (
	"log"
	"strconv"
)

func PartOne(input []string) string {
	max := 0
	curr := 0

	for _, line := range input {
		if line == "" {
			if curr > max {
				max = curr
			}
			curr = 0

			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Println(err)

			return ""
		}

		curr += cals
	}

	return strconv.Itoa(max)
}

func PartTwo(input []string) string {
	elves := make([]int, 0)
	curr := 0

	for i, line := range input {
		if i != 0 && line == "" {
			elves = append(elves, curr)
			curr = 0

			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Println(err)

			return ""
		}

		curr += cals
	}

	res := 0
	for i := len(elves) - 3; i < len(elves); i++ {
		res += elves[i]
	}

	return strconv.Itoa(res)
}
