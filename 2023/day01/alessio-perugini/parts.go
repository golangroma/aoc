package main

import (
	"fmt"
	"strconv"
)

func PartOne(input []string) string {
	var result int64
	for _, v := range input {
		calibrationvalue := ""

		for i := 0; i < len(v); i++ {
			if v[i] >= '0' && v[i] <= '9' {
				calibrationvalue = string(v[i])
				break
			}
		}
		for i := len(v) - 1; i >= 0; i-- {
			if v[i] >= '0' && v[i] <= '9' {
				calibrationvalue += string(v[i])
				break
			}
		}
		if calibrationvalue == "" {
			continue
		}
		cValue, err := strconv.ParseInt(calibrationvalue, 10, 64)
		if err != nil {
			panic(err)
		}
		result += cValue
	}
	return fmt.Sprintf("%v", result)
}

func PartTwo(input []string) string {
	var result int64
	for _, v := range input {
		calibrationvalue := ""
		for i := 0; i < len(v); i++ {
			if v[i] >= '0' && v[i] <= '9' {
				calibrationvalue = string(v[i])
				if candidate := containsNumberLeft(v[:i]); candidate != "" {
					calibrationvalue = candidate
				}
				break
			}
		}

		for i := len(v) - 1; i >= 0; i-- {
			if v[i] >= '0' && v[i] <= '9' {
				tmp := string(v[i])
				if candidate := containsNumberRight(v[i+1:]); candidate != "" {
					tmp = candidate
				}
				calibrationvalue += tmp
				break
			}
		}
		if calibrationvalue == "" {
			continue
		}
		cValue, err := strconv.ParseInt(calibrationvalue, 10, 64)
		if err != nil {
			panic(err)
		}
		result += cValue
	}
	return fmt.Sprintf("%v", result)
}

var numbersAsString = []string{"one", "two", "six", "nine", "four", "five", "three", "seven", "eight"}
var mapNumbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// 3, 4, 5
func containsNumberLeft(input string) string {
	if len(input) < 3 {
		return ""
	}
	for i := 0; i <= len(input)-3; i++ {
		for _, v := range []int{3, 4, 5} {
			if len(input)-i >= v {
				candidate := input[i : i+v]
				if number, ok := mapNumbers[candidate]; ok {
					return number
				}
			}
		}
	}

	return ""
}

func containsNumberRight(input string) string {
	if len(input) < 3 {
		return ""
	}
	for i := len(input) - 3; i >= 0; i-- {
		for _, v := range []int{3, 4, 5} {
			if len(input)-i >= v {
				candidate := input[i : i+v]
				if number, ok := mapNumbers[candidate]; ok {
					return number
				}
			}
		}
	}

	return ""
}
