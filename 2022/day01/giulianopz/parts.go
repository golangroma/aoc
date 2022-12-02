package main

import (
	"log"
	"strconv"
)

var maxCaloriesPerElf uint64
var topThree = []uint64{0, 0, 0}

func PartOne(input []string) string {

	var caloriesPerElf uint64

	for _, line := range input {
		if line == "" {
			if caloriesPerElf > maxCaloriesPerElf {
				maxCaloriesPerElf = caloriesPerElf
			}
			caloriesPerElf = 0
			continue
		}

		caloriesPerElf += toUint(line)
	}
	return strconv.FormatUint(maxCaloriesPerElf, 10)
}

func PartTwo(input []string) string {

	var caloriesPerElf uint64

	for _, line := range input {
		if line == "" {
			for idx, top := range topThree {
				if caloriesPerElf > top {

					topThree[idx] = caloriesPerElf

					if idx == 0 {
						tmp := topThree[idx+1]
						topThree[idx+1] = top
						topThree[idx+2] = tmp
					}

					if idx == 1 {
						topThree[idx+1] = top
					}
					break
				}
			}
			caloriesPerElf = 0
			continue
		}

		caloriesPerElf += toUint(line)
	}
	return strconv.FormatUint(topThree[0]+topThree[1]+topThree[2], 10)
}

func toUint(s string) uint64 {
	parsed, err := strconv.ParseUint(s, 10, 64)
	exitFatally(err)
	return parsed
}

func exitFatally(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
