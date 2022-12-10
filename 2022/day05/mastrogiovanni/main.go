package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golangroma/aoc/utils"
)

func ReadFileNoTrim(filename string) []string {
	content, err := os.ReadFile(filename)
	utils.CheckErr(err)

	return SplitInputNoTrim(string(content))
}

func SplitInputNoTrim(content string) []string {
	stringContent := string(content)
	return strings.Split(stringContent, "\n")
}

func main() {
	input := ReadFileNoTrim("input.txt")

	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}
