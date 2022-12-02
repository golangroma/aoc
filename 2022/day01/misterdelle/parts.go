package main

import (
	"strconv"
)

func PartOne(input []string) string {
	var elvesMap = make(map[int]int)
	var elfId int = 1

	for _, v := range input {
		if v != "" {
			calories, _ := strconv.Atoi(v)
			elvesMap[elfId] = elvesMap[elfId] + calories
		} else {
			// Blank line means new elf
			elfId++
		}
	}

	_, maxCalories := getMaxCaloriesXElf(elvesMap)

	return strconv.Itoa(maxCalories)
}

func PartTwo(input []string) string {
	var elvesMap = make(map[int]int)
	var elfId int = 1

	for _, v := range input {
		if v != "" {
			calories, _ := strconv.Atoi(v)
			elvesMap[elfId] = elvesMap[elfId] + calories
		} else {
			// Blank line means new elf
			elfId++
		}
	}

	elfCarryingMaxCalories1, maxCalories1 := getMaxCaloriesXElf(elvesMap)
	elvesMap[elfCarryingMaxCalories1] = 0
	elfCarryingMaxCalories2, maxCalories2 := getMaxCaloriesXElf(elvesMap)
	elvesMap[elfCarryingMaxCalories2] = 0
	elfCarryingMaxCalories3, maxCalories3 := getMaxCaloriesXElf(elvesMap)
	elvesMap[elfCarryingMaxCalories3] = 0

	maxCalories := maxCalories1 + maxCalories2 + maxCalories3

	return strconv.Itoa(maxCalories)
}

func getMaxCaloriesXElf(elvesMap map[int]int) (int, int) {
	elfCarryingMaxCalories := 0
	maxCalories := 0

	for k, v := range elvesMap {
		if v > maxCalories {
			maxCalories = v
			elfCarryingMaxCalories = k
		}
	}

	return elfCarryingMaxCalories, maxCalories
}
