package main

import (
	"strconv"
	"strings"
	"unicode"
)

func PartOne(input []string) string {
	var tot int

	for _, line := range input {
		var result string

		runes := []rune(line)
		for _, r := range runes {
			if unicode.IsDigit(r) {
				result = string(r)
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				result += string(runes[i])
				break
			}
		}

		res, err := strconv.Atoi(result)
		if err != nil {
			panic(err)
		}
		tot += res
	}

	return strconv.Itoa(tot)
}

func PartTwo(input []string) string {
	tot := 0

	for _, line := range input {
		digits := []string{}

		for i, r := range line {
			if unicode.IsDigit(r) {
				digits = append(digits, string(r))
			} else {
				if digitStr, found := hasDigitPrefix(line[i:]); found {
					digits = append(digits, digitStr)
				}
			}
		}

		concat := digits[0] + digits[len(digits)-1]

		result, err := strconv.Atoi(concat)
		if err != nil {
			panic(err)
		}
		tot += result
	}

	return strconv.Itoa(tot)
}

func hasDigitPrefix(s string) (string, bool) {
	digits := []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	for i, digitString := range digits {
		if strings.HasPrefix(s, digitString) {
			return strconv.Itoa(i), true
		}
	}
	return "", false
}
