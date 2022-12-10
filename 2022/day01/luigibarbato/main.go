package main

import (
	"fmt"

	"github.com/golangroma/aoc/utils"
)

func main() {
	input := utils.ReadFile("input.txt")
	r, err := PartOne(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Solution for part one:", r)

	r, err = PartTwo(input)
	if err != nil {
		panic(err)
	}

	fmt.Println("Solution for part two:", r)
}
