package main

import (
	"fmt"

	"github.com/golangroma/aoc/utils"
)

func main() {
	input := utils.ReadFile("input.txt")

	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
	/*
		Part 1: 13809
		Part 2: 12316
	*/
}
