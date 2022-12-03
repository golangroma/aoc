package main

import (
	"fmt"
	"sort"
	"strconv"
)

func PartOne(input []string) string {
	current_elf_calories := 0
	max_calories := 0
	for _, item := range input {
		if item == "" {
			if current_elf_calories > max_calories {
				max_calories = current_elf_calories
			}
			current_elf_calories = 0
		} else {
			value, _ := strconv.Atoi(item)
			current_elf_calories += value
		}
	}
	if current_elf_calories > max_calories {
		max_calories = current_elf_calories
	}
	return fmt.Sprintf("%d", max_calories)
}

func PartTwo(input []string) string {
	current_elf_calories := 0
	elves := make([]int, 0)
	for _, item := range input {
		if item == "" {
			elves = append(elves, current_elf_calories)
			current_elf_calories = 0
		} else {
			value, _ := strconv.Atoi(item)
			current_elf_calories += value
		}
	}

	elves = append(elves, current_elf_calories)
	sort.Ints(elves)

	tot := 0
	for i := len(elves) - 1 - 2; i < len(elves); i++ {
		tot += elves[i]
	}

	return fmt.Sprintf("%d", tot)
}
