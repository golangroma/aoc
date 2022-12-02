package main

import (
	"fmt"
	"strings"
)

const (
	NoChoice HandShape = iota
	Rock
	Paper
	Scissor
)

type HandShape int

const (
	Lose Points = 0
	Draw Points = 3
	Win  Points = 6
)

type Points int

func PartOne(input []string) string {
	totalScore := 0

	for _, v := range input {
		matchPredictions := strings.Split(v, " ")
		opponentChoise := parseChoise(matchPredictions[0])
		prediction := parseChoise(matchPredictions[1])

		totalScore += round(opponentChoise, prediction)

	}

	return fmt.Sprintf("%d", totalScore)
}

func parseChoise(c string) HandShape {
	switch c {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissor
	}

	return NoChoice
}

func round(opponent, me HandShape) int {
	if opponent == me { //draw
		return int(Draw) + int(me)
	}

	matchResult := int(Lose)

	switch true {
	case opponent == Rock && me == Paper:
		matchResult = int(Win)
	case opponent == Paper && me == Scissor:
		matchResult = int(Win)
	case opponent == Scissor && me == Rock:
		matchResult = int(Win)
	}

	return matchResult + int(me)
}

func parseChoiseStrategically(opponent HandShape, strategy string) HandShape {
	if strategy == "Y" { // we need to draw
		return opponent
	}

	// we need to win
	if strategy == "Z" {
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissor
		case Scissor:
			return Rock
		}
	}

	// we need to lose
	if strategy == "X" {
		switch opponent {
		case Rock:
			return Scissor
		case Paper:
			return Rock
		case Scissor:
			return Paper
		}
	}

	return NoChoice
}

func PartTwo(input []string) string {
	totalScore := 0

	for _, v := range input {
		matchPredictions := strings.Split(v, " ")
		opponentChoise := parseChoise(matchPredictions[0])
		prediction := parseChoiseStrategically(opponentChoise, matchPredictions[1])

		totalScore += round(opponentChoise, prediction)
	}

	return fmt.Sprintf("%d", totalScore)
}
