package main

import "fmt"

func PartOne(input []string) string {
	sum := 0

	for _, rucksack := range input {
		sum += getPriorityOfCommonItems(rucksack)
	}

	return fmt.Sprintf("%d", sum)
}

func getPriorityOfCommonItems(rucksack string) int {
	sum := 0
	slot1, itemsInCommon := make(map[rune]int), make(map[rune]bool)

	l := len(rucksack)
	for i, item := range rucksack {
		if i < l/2 {
			slot1[item]++
			continue
		}
		if _, ok := slot1[item]; ok && !itemsInCommon[item] {
			itemsInCommon[item] = true
			sum += itemPriority(item)
		}
	}

	return sum
}

func itemPriority(item rune) int {
	lowerCaseOffset, upperCaseOffset := int('a'), int('A')

	if item >= 'a' && item <= 'z' {
		return int(item) - lowerCaseOffset + 1
	}
	if item >= 'A' && item <= 'Z' {
		return int(item) - upperCaseOffset + 27
	}

	return 0
}

func PartTwo(input []string) string {
	sum := 0

	for i := 0; i < len(input); i += 3 {
		sum += commonItemPriorityBetweenThreeElves(input[i : i+3])
	}

	return fmt.Sprintf("%d", sum)
}

func commonItemPriorityBetweenThreeElves(rucksacks []string) int {
	sum := 0
	lastRucksack := len(rucksacks) - 1
	common := make(map[rune]int)

	for i, rucksack := range rucksacks {
		for _, item := range rucksack {
			if common[item] != i {
				continue
			}
			common[item]++
			if lastRucksack == i {
				sum += itemPriority(item)
			}
		}
	}

	return sum
}
