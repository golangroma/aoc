package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/golangroma/aoc/utils"
)

func main() {
	content, err := os.ReadFile("input.txt")
	utils.CheckErr(err)
	lines := strings.Split(string(content), "\n")

	fmt.Printf("Part 1: %v\n", PartOne(lines))
	fmt.Printf("Part 2: %v\n", PartTwo(lines))
}
