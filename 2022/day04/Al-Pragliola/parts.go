package main

import (
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	totalContained := 0

	for _, line := range input {
		ranges := strings.Split(line, ",")
		rangeOne := strings.Split(ranges[0], "-")
		rangeTwo := strings.Split(ranges[1], "-")

		rangeOneStart, _ := strconv.Atoi(rangeOne[0])
		rangeOneEnd, _ := strconv.Atoi(rangeOne[1])
		rangeTwoStart, _ := strconv.Atoi(rangeTwo[0])
		rangeTwoEnd, _ := strconv.Atoi(rangeTwo[1])

		if rangeOneStart <= rangeTwoStart && rangeOneEnd >= rangeTwoEnd {
			totalContained++
			continue
		}

		if rangeTwoStart <= rangeOneStart && rangeTwoEnd >= rangeOneEnd {
			totalContained++
			continue
		}
	}

	return strconv.Itoa(totalContained)
}

func PartTwo(input []string) string {
	totalOverlaps := 0

	for _, line := range input {
		ranges := strings.Split(line, ",")
		rangeOne := strings.Split(ranges[0], "-")
		rangeTwo := strings.Split(ranges[1], "-")

		rangeOneStart, _ := strconv.Atoi(rangeOne[0])
		rangeOneEnd, _ := strconv.Atoi(rangeOne[1])
		rangeTwoStart, _ := strconv.Atoi(rangeTwo[0])
		rangeTwoEnd, _ := strconv.Atoi(rangeTwo[1])

		if (rangeOneStart < rangeTwoStart && rangeOneEnd < rangeTwoStart) ||
			(rangeOneStart > rangeTwoEnd && rangeOneEnd > rangeTwoEnd) {
			continue
		}

		if (rangeTwoStart < rangeOneStart && rangeTwoEnd < rangeOneStart) ||
			(rangeTwoStart > rangeOneEnd && rangeTwoEnd > rangeOneEnd) {
			continue
		}

		totalOverlaps++
	}

	return strconv.Itoa(totalOverlaps)
}
