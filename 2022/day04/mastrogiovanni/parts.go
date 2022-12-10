package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CheckIfContains(a, b, c, d int) bool {
	return a <= c && b >= d || c <= a && d >= b
}

func CheckIfOverlaps(a, b, c, d int) bool {
	if c >= a && c <= b || d >= a && d <= b {
		return true
	}
	if a >= c && a <= d || b >= c && b <= d {
		return true
	}
	return false
}

func PartOne(input []string) string {
	tot := 0
	for _, row := range input {
		ranges := strings.Split(row, ",")
		pair1 := strings.Split(ranges[0], "-")
		pair2 := strings.Split(ranges[1], "-")
		a, _ := strconv.Atoi(pair1[0])
		b, _ := strconv.Atoi(pair1[1])
		c, _ := strconv.Atoi(pair2[0])
		d, _ := strconv.Atoi(pair2[1])
		contains := CheckIfContains(a, b, c, d)
		if contains {
			tot += 1
		}
	}
	return fmt.Sprintf("%d", tot)
}

func PartTwo(input []string) string {
	tot := 0
	for _, row := range input {
		ranges := strings.Split(row, ",")
		pair1 := strings.Split(ranges[0], "-")
		pair2 := strings.Split(ranges[1], "-")
		a, _ := strconv.Atoi(pair1[0])
		b, _ := strconv.Atoi(pair1[1])
		c, _ := strconv.Atoi(pair2[0])
		d, _ := strconv.Atoi(pair2[1])
		contains := CheckIfOverlaps(a, b, c, d)
		if contains {
			tot += 1
		}
	}
	return fmt.Sprintf("%d", tot)
}
