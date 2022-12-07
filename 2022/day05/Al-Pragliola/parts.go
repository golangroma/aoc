package main

import (
	"strconv"
	"strings"
)

func identity(s string) string {
	return s
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func parseStacks(input []string) ([]string, int) {
	breakingIndex := 0

	for i, line := range input {
		if line == "" {
			breakingIndex = i
			break
		}
	}

	stacksNum := (len(input[breakingIndex-1]) + 1) / 4

	stacks := make([]string, stacksNum)

	for j := breakingIndex - 2; j >= 0; j-- {
		for i := 0; i < stacksNum; i++ {
			newItemIdx := i*4 + 1

			if newItemIdx >= len(input[j]) {
				break
			}

			if input[j][newItemIdx] == ' ' {
				continue
			}

			stacks[i] = string(input[j][newItemIdx]) + stacks[i]
		}
	}

	return stacks, breakingIndex
}

func doInstructions(input []string, breakingIndex int, stacks []string, strategy func(string) string) []string {
	for i := breakingIndex + 1; i < len(input); i++ {
		if input[i] == "" {
			break
		}

		instr := strings.Split(input[i], " ")

		num, _ := strconv.Atoi(instr[1])
		from, _ := strconv.Atoi(instr[3])
		to, _ := strconv.Atoi(instr[5])

		strFrom := stacks[from-1][:num]
		stacks[from-1] = stacks[from-1][num:]
		stacks[to-1] = strategy(strFrom) + stacks[to-1]
	}

	return stacks
}

func PartOne(input []string) string {
	result := ""

	stacks, breakingIndex := parseStacks(input)

	stacks = doInstructions(input, breakingIndex, stacks, reverseString)

	for _, s := range stacks {
		if len(s) > 0 {
			result += string(s[0])
		}
	}

	return result
}

func PartTwo(input []string) string {
	result := ""

	stacks, breakingIndex := parseStacks(input)

	stacks = doInstructions(input, breakingIndex, stacks, identity)

	for _, s := range stacks {
		if len(s) > 0 {
			result += string(s[0])
		}
	}

	return result
}
