package main

import (
	"strconv"
)

func PartOne(input []string) string {
	signal := input[0]

	for i := 0; i < len(signal)-3; i++ {
		s := signal[i : i+4]
		if areUniqueChars(s) {
			return strconv.Itoa(i + 4)
		}
	}

	return ""
}

func areUniqueChars(s string) bool {
	unique := make(map[rune]struct{})
	for _, r := range s {
		unique[r] = struct{}{}
	}
	return len(unique) == len(s)
}

func PartTwo(input []string) string {
	signal := input[0]

	for i := 0; i < len(signal)-13; i++ {
		s := signal[i : i+14]
		if areUniqueChars(s) {
			return strconv.Itoa(i + 14)
		}
	}

	return ""
}
