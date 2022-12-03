package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

var lowercase []string = make([]string, 0)
var uppercase []string = make([]string, 0)

func init() {
	for r := 'a'; r <= 'z'; r++ {
		lowercase = append(lowercase, fmt.Sprintf("%c", r))
		R := unicode.ToUpper(r)
		uppercase = append(uppercase, fmt.Sprintf("%c", R))
	}
}

func PartOne(input []string) string {

	var totalPriorities int
	for _, line := range input {

		separator := len(line) / 2
		firstHalf := strings.Split(line[:separator], "")
		secondHalf := line[separator:]

		sort.Strings(firstHalf)
		for i, c := range firstHalf {
			if i == 0 || (firstHalf[i] != firstHalf[i-1]) {
				if strings.Contains(secondHalf, c) {
					prior := priority(c)
					totalPriorities += prior
				}
			}
		}
	}
	return fmt.Sprintf("%d", totalPriorities)
}

const offset = 1

func priority(s string) int {
	for i, c := range lowercase {
		if c == s {
			return i + offset
		}
	}
	for i, c := range uppercase {
		if c == s {
			return i + len(lowercase) + offset
		}
	}
	return 0
}

func PartTwo(input []string) string {
	return ""
}
