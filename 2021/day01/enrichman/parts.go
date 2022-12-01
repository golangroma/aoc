package main

import (
	"strconv"

	"github.com/golangroma/aoc/cli/utils"
)

func PartOne(input []string) string {
	measurements, err := utils.Convert(input, utils.StringSliceToIntSliceConverter)
	utils.CheckErr(err)

	count := 0

	// start from the second
	for i := 1; i < len(measurements); i++ {
		prev := measurements[i-1]
		curr := measurements[i]
		if prev < curr {
			count++
		}
	}

	return strconv.Itoa(count)
}

func PartTwo(input []string) string {
	measurements, err := utils.Convert(input, utils.StringSliceToIntSliceConverter)
	utils.CheckErr(err)

	count := 0

	// start from the second window
	for i := 3; i < len(measurements); i++ {
		prevWindow := measurements[i-3] + measurements[i-2] + measurements[i-1]
		currWindow := measurements[i-2] + measurements[i-1] + measurements[i]
		if prevWindow < currWindow {
			count++
		}
	}

	return strconv.Itoa(count)
}
