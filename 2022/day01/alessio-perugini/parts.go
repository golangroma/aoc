package main

import (
	"fmt"
	"strconv"
)

func PartOne(input []string) string {
	mostCalories, tmp := 0, 0

	for _, calorie := range input {
		if calorie == "" {
			if tmp > mostCalories {
				mostCalories = tmp
			}
			tmp = 0
			continue
		}
		v, _ := strconv.Atoi(calorie)
		tmp += v
	}

	if tmp > 0 && tmp > mostCalories {
		mostCalories = tmp
	}

	return fmt.Sprintf("%d", mostCalories)
}

func PartTwo(input []string) string {
	sumTopCalories, tmp := 0, 0
	top3calories := make([]int, 3)

	for _, calorie := range input {
		if calorie == "" {
			top3calories = insertOrderedSlice(top3calories, tmp)
			tmp = 0
			continue
		}
		v, _ := strconv.Atoi(calorie)
		tmp += v
	}
	top3calories = insertOrderedSlice(top3calories, tmp)

	for _, v := range top3calories {
		sumTopCalories += v
	}

	return fmt.Sprintf("%d", sumTopCalories)
}

func insertOrderedSlice(calories []int, calorie int) []int {
	for i := 0; i < len(calories); i++ {
		if calorie <= calories[i] {
			continue
		}

		copy(calories[i+1:], calories[i:])
		calories[i] = calorie
		break
	}

	return calories[:3]
}
