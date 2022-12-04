package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func PartOne(input []string) string {

	var fullyContainingAssignmentPairs int
	for _, line := range input {
		pair := strings.Split(line, ",")
		firstElf := strings.Split(pair[0], "-")
		secondElf := strings.Split(pair[1], "-")

		if (asInt(firstElf[0]) <= asInt(secondElf[0]) && asInt(firstElf[1]) >= asInt(secondElf[1])) ||
			(asInt(secondElf[0]) <= asInt(firstElf[0]) && asInt(secondElf[1]) >= asInt(firstElf[1])) {
			fullyContainingAssignmentPairs += 1
		}
	}
	return fmt.Sprintf("%d", fullyContainingAssignmentPairs)
}

func PartTwo(input []string) string {

	var overlappingAssignmentPairs int
	for _, line := range input {
		pair := strings.Split(line, ",")
		firstElf := strings.Split(pair[0], "-")
		secondElf := strings.Split(pair[1], "-")

		if (asInt(firstElf[0]) <= asInt(secondElf[0]) && asInt(firstElf[1]) >= asInt(secondElf[1])) ||
			(asInt(secondElf[0]) <= asInt(firstElf[0]) && asInt(secondElf[1]) >= asInt(firstElf[1]) ||
				asInt(firstElf[0]) <= asInt(secondElf[1]) && asInt(firstElf[1]) >= asInt(secondElf[0])) {
			overlappingAssignmentPairs += 1
		}
	}
	return fmt.Sprintf("%d", overlappingAssignmentPairs)
}

func asInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
