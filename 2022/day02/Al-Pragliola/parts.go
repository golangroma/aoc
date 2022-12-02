package main

import (
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	score := 0

	for _, line := range input {
		moves := strings.Split(line, " ")
		moveOne := moves[0]
		moveTwo := moves[1]

		score += basicStrat(moveOne, moveTwo)
	}

	return strconv.Itoa(score)
}

func PartTwo(input []string) string {
	score := 0

	for _, line := range input {
		moves := strings.Split(line, " ")
		moveOne := moves[0]
		result := moves[1]

		score += complexStrat(moveOne, result)
	}

	return strconv.Itoa(score)
}

func complexStrat(moveOne, result string) int {
	resultScores := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	mapIndex := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
		"A": 0,
		"B": 1,
		"C": 2,
	}

	outMat := [][]int{
		{3, 1, 2},
		{1, 2, 3},
		{2, 3, 1},
	}

	return outMat[mapIndex[moveOne]][mapIndex[result]] + resultScores[result]

}

func basicStrat(moveOne, moveTwo string) int {
	moveScores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	return moveScores[moveTwo] + outcome(moveOne, moveTwo)

}

func outcome(moveOne, moveTwo string) int {
	win := 6
	lose := 0
	tie := 3

	mapIndex := map[string]int{
		"X": 0,
		"Y": 1,
		"Z": 2,
		"A": 0,
		"B": 1,
		"C": 2,
	}

	outMat := [][]int{
		{0, 1, -1},
		{-1, 0, 1},
		{1, -1, 0},
	}

	if outMat[mapIndex[moveOne]][mapIndex[moveTwo]] == 0 {
		return tie
	}

	if outMat[mapIndex[moveOne]][mapIndex[moveTwo]] == 1 {
		return win
	}

	return lose
}
