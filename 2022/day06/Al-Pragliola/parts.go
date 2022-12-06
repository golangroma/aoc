package main

import "strconv"

func PartOne(input []string) string {
	msg := input[0]

	return strconv.Itoa(calcMsgIdx(msg, 4))
}

func PartTwo(input []string) string {
	msg := input[0]

	return strconv.Itoa(calcMsgIdx(msg, 14))
}

func calcMsgIdx(msg string, num int) int {
	diffCharsNum := 0

	for i := 0; i < len(msg); i++ {
		diffChars := make(map[rune]int)

		for j := i; j < i+num; j++ {
			if diffChars[rune(msg[j])] == 0 {
				diffCharsNum++
			}

			diffChars[rune(msg[j])]++
		}

		if diffCharsNum == num {
			return i + num
		}

		diffCharsNum = 0
	}

	return 0
}
