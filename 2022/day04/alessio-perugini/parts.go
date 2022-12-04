package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	cnt := 0

	for _, v := range input {
		pairs := strings.Split(v, ",")

		min1, max1 := parsePairRange(pairs[0])
		min2, max2 := parsePairRange(pairs[1])

		if fullyContainedAinB(min1, max1, min2, max2) || fullyContainedAinB(min2, max2, min1, max1) {
			cnt++
		}
	}

	return fmt.Sprintf("%d", cnt)
}

func fullyContainedAinB(minA, maxA, minB, maxB int) bool {
	return minA <= minB && maxA >= maxB
}

func parsePairRange(pair string) (int, int) {
	interval := strings.Split(pair, "-")
	min, _ := strconv.Atoi(interval[0])
	max, _ := strconv.Atoi(interval[1])
	return min, max
}

func PartTwo(input []string) string {
	cnt := 0

	for _, v := range input {
		pairs := strings.Split(v, ",")

		min1, max1 := parsePairRange(pairs[0])
		min2, max2 := parsePairRange(pairs[1])

		if (min1 >= min2 && min1 <= max2) ||
			(min2 >= min1 && min2 <= max1) ||
			(max1 >= min2 && max1 <= max2) {
			cnt++
		}
	}

	return fmt.Sprintf("%d", cnt)
}
