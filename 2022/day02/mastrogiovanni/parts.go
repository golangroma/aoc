package main

import (
	"fmt"
	"strings"
)

const (
	Rock int = iota
	Paper
	Scissors
)

const (
	PlayerA int = iota
	PlayerB
	Draw
)

func convert(item string) int {
	switch item {
	case "A":
		return Rock // rock
	case "B":
		return Paper // paper
	case "C":
		return Scissors // scissors
	case "X":
		return Rock // rock
	case "Y":
		return Paper // paper
	case "Z":
		return Scissors // scissors
	}
	panic(fmt.Errorf("error in parsing %s", item))
}

func score(item int) int {
	switch item {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic(fmt.Errorf("error in assigning score %d", item))
}

func winner(moveA int, moveB int) int {
	if moveA == moveB {
		return Draw
	}
	if moveA == Rock && moveB == Scissors {
		return PlayerA
	}
	if moveA == Scissors && moveB == Paper {
		return PlayerA
	}
	if moveA == Paper && moveB == Rock {
		return PlayerA
	}
	return PlayerB
}

func moveBFromMoveAAndOutcome(moveA int, outcome string) int {
	// B need to loose
	if outcome == "X" {
		switch moveA {
		case Rock:
			return Scissors
		case Scissors:
			return Paper
		case Paper:
			return Rock
		}
	}
	// B need to draw
	if outcome == "Y" {
		return moveA
	}
	// B need to win
	if outcome == "Z" {
		switch moveA {
		case Rock:
			return Paper
		case Scissors:
			return Rock
		case Paper:
			return Scissors
		}
	}
	panic(fmt.Errorf("error in parsing outcome %s or move %d", outcome, moveA))
}

func PartOne(input []string) string {
	total := 0
	for _, row := range input {
		components := strings.Split(row, " ")
		moveA := convert(components[0])
		moveB := convert(components[1])
		result := winner(moveA, moveB)
		total += score(moveB)
		if result == Draw {
			total += 3
		} else if result == PlayerB {
			total += 6
		}
	}
	return fmt.Sprintf("%d", total)
}

func PartTwo(input []string) string {
	total := 0
	for _, row := range input {
		components := strings.Split(row, " ")
		moveA := convert(components[0])
		outcome := components[1]
		moveB := moveBFromMoveAAndOutcome(moveA, outcome)
		result := winner(moveA, moveB)
		total += score(moveB)
		if result == Draw {
			total += 3
		} else if result == PlayerB {
			total += 6
		}
	}
	return fmt.Sprintf("%d", total)
}
