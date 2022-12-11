package main

import (
	"strconv"
	"unicode"
)

func PartOne(input []string) string {
	priorities := buildPriorities()
	commonTotal := 0

	for _, line := range input {
		common := 0

		rucksack := make(map[rune]int)
		errors := make(map[rune]int)

		for r := range line[:len(line)/2] {
			rucksack[rune(line[r])]++
		}

		for _, s := range line[len(line)/2:] {
			if rucksack[s] > 0 && errors[s] == 0 {
				errors[s]++
				common += priorities[s]
			}
		}

		commonTotal += common
	}

	return strconv.Itoa(commonTotal)
}

func PartTwo(input []string) string {
	priorities := buildPriorities()
	commonTotal := 0

	for i := 0; i < len(input); i += 3 {
		common := 0

		rucksackOne := make(map[rune]int)
		rucksackTwo := make(map[rune]int)

		for r := range input[i] {
			rucksackOne[rune(input[i][r])]++
		}

		for r := range input[i+1] {
			if rucksackOne[rune(input[i+1][r])] > 0 {
				rucksackTwo[rune(input[i+1][r])]++
			}
		}

		for r := range input[i+2] {
			if rucksackTwo[rune(input[i+2][r])] > 0 {
				common += priorities[rune(input[i+2][r])]
				break
			}
		}

		commonTotal += common
	}

	return strconv.Itoa(commonTotal)
}

func buildPriorities() map[rune]int {
	priorities := make(map[rune]int)

	for i, r := range "abcdefghijklmnopqrstuvwxyz" {
		priorities[r] = i + 1
		priorities[unicode.ToUpper(r)] = i + 1 + 26
	}

	return priorities
}
