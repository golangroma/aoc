package main

import (
	"fmt"

	"github.com/golangroma/aoc/utils"
)

func main() {
	input := utils.ReadFile("input.txt")

	fmt.Printf("Part 1: %v\n", PartOne(input, 2000000))
	fmt.Printf("Part 2: %v\n", PartTwo(input, 4000000))
}
