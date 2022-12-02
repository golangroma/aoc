package main

import (
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	// Opponent (first element of slice)
	// A for Rock, B for Paper, and C for Scissors
	//
	// Me (second element of slice)
	// X for Rock, Y for Paper, and Z for Scissors
	//
	// Points for shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
	// Outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
	//

	totalPoints := 0

	totalRulesMap := map[string]int{
		"A X": 4, // Rock/Rock/Draw
		"A Y": 8, // Rock/Paper/Win
		"A Z": 3, // Rock/Scissors/Lose
		"B X": 1, // Paper/Rock/Lose
		"B Y": 5, // Paper/Paper/Draw
		"B Z": 9, // Paper/Scissors/Win
		"C X": 7, // Scissors/Rock/Win
		"C Y": 2, // Scissors/Paper/Lose
		"C Z": 6, // Scissors/Scissors/Draw
	}

	for _, v := range input {
		totalPoints += totalRulesMap[v]
	}

	return strconv.Itoa(totalPoints)
}

func PartTwo(input []string) string {
	//
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	//
	totalPoints := 0

	for _, v := range input {

		roundSelections := strings.Split(v, " ")
		opponentSelection := roundSelections[0]
		mineSelection := roundSelections[1]

		totalPoints += getHandsetStrategically(opponentSelection, mineSelection)
	}

	return strconv.Itoa(totalPoints)
}

func getHandsetStrategically(opponent, me string) int {
	//
	// Opponent (first element of slice)
	// A for Rock, B for Paper, and C for Scissors
	//
	// Me (second element of slice)
	// X for Rock, Y for Paper, and Z for Scissors
	//
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
	//

	rc := 0

	switch opponent {
	case "A":
		// "A X": 4, // Rock/Rock/Draw
		// "A Y": 8, // Rock/Paper/Win
		// "A Z": 3, // Rock/Scissors/Lose
		if me == "X" {
			// Rock/Scissors
			// from Draw to Lose
			// "A Z": 3, // Rock/Scissors/Lose
			rc = 3
		} else if me == "Y" {
			// Rock/Paper
			// from Win to Draw
			// "A X": 4, // Rock/Rock/Draw
			rc = 4
		} else if me == "Z" {
			// Rock/Scissors/Lose
			// From Lose to Win
			// "A Y": 8, // Rock/Paper/Win
			rc = 8
		}
	case "B":
		// "B X": 1, // Paper/Rock/Lose
		// "B Y": 5, // Paper/Paper/Draw
		// "B Z": 9, // Paper/Scissors/Win
		if me == "X" {
			// Paper/Rock
			// from Lose to Lose
			// "B X": 1, // Paper/Rock/Lose
			rc = 1
		} else if me == "Y" {
			// Paper/Paper
			// from Draw to Draw
			// "B Y": 5, // Paper/Paper/Draw
			rc = 5
		} else if me == "Z" {
			// Paper/Scissors
			// From Win to Win
			// "B Z": 9, // Paper/Scissors/Win
			rc = 9
		}
	case "C":
		// "C X": 7, // Scissors/Rock/Win
		// "C Y": 2, // Scissors/Paper/Lose
		// "C Z": 6, // Scissors/Scissors/Draw
		if me == "X" {
			// Scissors/Rock
			// From Win to Lose
			// "C Y": 2, // Scissors/Paper/Lose
			rc = 2
		} else if me == "Y" {
			// Scissors/Paper
			// From Lose to Draw
			// "C Z": 6, // Scissors/Scissors/Draw
			rc = 6
		} else if me == "Z" {
			// "C Z": 6, // Scissors/Scissors
			// From Draw to Win
			// "C X": 7, // Scissors/Rock/Win
			rc = 7
		}
	}

	return rc
}
