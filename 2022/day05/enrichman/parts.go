package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golangroma/aoc/utils"
)

type Stacks []Stack

type Stack []Crate

type Crate string

func PartOne(input []string) string {
	parseInput(input)
	return ""
}

func PartTwo(input []string) string {
	return ""
}

func parseInput(input []string) {
	stackInput := []string{}
	//arrangementsInput := []string{}

	for i, line := range input {
		if line == "" {
			stackInput = input[0:i]
			//arrangementsInput = input[i+1:]
			break
		}
	}

	//fmt.Println(stackInput, arrangementsInput)

	printDrawing(stackInput)

	var stacks Stacks
	for i := len(stackInput) - 1; i >= 0; i-- {
		if i == len(stackInput)-1 {
			stacksIDs := strings.Fields(stackInput[i])
			lastStackIDStr := stacksIDs[len(stacksIDs)-1]
			lastStackID, err := strconv.Atoi(lastStackIDStr)
			utils.CheckErr(err)

			stacks = make(Stacks, lastStackID)
			continue
		}

		cratesLine := strings.Fields(stackInput[i])
		fmt.Println(cratesLine)
		for _, crate := range cratesLine {
			if stacks[i] == nil {
				stacks[i] = Stack{}
			}
			stack := stacks[i]
			stack = append(stack, Crate(crate))
			stacks[i] = stack
		}

		fmt.Println(stacks[1])
	}
	fmt.Println(len(stacks))
}

func printDrawing(drawing []string) {
	for i, line := range drawing {
		if strings.HasPrefix(line, "    ") {
			line = strings.Replace(line, "    ", "[-] ", 1)
		}
		line = strings.ReplaceAll(line, "    ", " [-]")

		drawing[i] = line
		fmt.Println(line)
	}
}
