package main

import (
	"fmt"
	"strconv"
	"strings"
)

var rules = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func PartOne(input []string) string {
	var result int
	for _, game := range input {
		gameID, sets, _ := strings.Cut(game, ": ")
		gameID = strings.TrimPrefix(gameID, "Game ")
		id := mustConvertToInt(gameID)
		valideGame := true
	SETS:
		for _, set := range strings.Split(sets, "; ") {
			for _, pair := range strings.Split(set, ", ") {
				n, color, _ := strings.Cut(pair, " ")
				if mustConvertToInt(n) > rules[color] {
					valideGame = false
					break SETS
				}
			}
		}
		if valideGame {
			result += id
		}
	}
	return fmt.Sprintf("%v", result)
}

func PartTwo(input []string) string {
	var result int
	for _, game := range input {
		_, sets, _ := strings.Cut(game, ": ")
		minimumSet := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}
		for _, set := range strings.Split(sets, "; ") {
			for _, pair := range strings.Split(set, ", ") {
				n, color, _ := strings.Cut(pair, " ")
				nInt := mustConvertToInt(n)
				if nInt > minimumSet[color] {
					minimumSet[color] = nInt
				}
			}
		}
		result += minimumSet["red"] * minimumSet["green"] * minimumSet["blue"]
	}
	return fmt.Sprintf("%v", result)
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
