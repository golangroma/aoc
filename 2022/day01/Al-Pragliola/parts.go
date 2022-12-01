package main

import (
	"sort"
	"strconv"
)

func PartOne(input []string) string {
	maxCal := 0
	cal := 0

	for _, v := range input {
		if v == "" {
			if cal > maxCal {
				maxCal = cal
			}

			cal = 0
			continue
		}

		currentCal, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		cal += currentCal
	}

	return strconv.Itoa(maxCal)
}

func PartTwo(input []string) string {
	calories := make([]int, 0)
	cal := 0

	for _, v := range input {
		if v == "" {
			calories = append(calories, cal)

			cal = 0
			continue
		}

		currentCal, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		cal += currentCal
	}

	sort.Ints(calories)

	return strconv.Itoa(calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3])
}
