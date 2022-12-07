package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golangroma/aoc/utils"
)

func main() {
	input := ReadFile("input.txt")

	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	utils.CheckErr(err)

	return SplitInput(string(content))
}

func SplitInput(content string) []string {
	return strings.Split(content, "\n")
}
