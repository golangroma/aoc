package main

import (
	"strconv"
	"strings"
)

const (
	rock     string = "Rock"
	paper    string = "Paper"
	scissors string = "Scissors"
)

type Shape struct {
	score int
	name  string
}

// A for Rock, B for Paper, and C for Scissors
var opponentDict map[string]*Shape = map[string]*Shape{"A": {1, rock}, "B": {2, paper}, "C": {3, scissors}}

// X for Rock, Y for Paper, and Z for Scissors
var myDict map[string]*Shape = map[string]*Shape{"X": {1, rock}, "Y": {2, paper}, "Z": {3, scissors}}

func reverseLookUp(shapeName string) string {
	for symbol, shape := range myDict {
		if shape.name == shapeName {
			return symbol
		}
	}
	return ""
}

var shapeOrderedSet []string = []string{rock, scissors, paper}

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
func won(myShapeName, opponentShapeName string) bool {
	return (myShapeName == rock && opponentShapeName == scissors) || (myShapeName == scissors && opponentShapeName == paper) || (myShapeName == paper && opponentShapeName == rock)
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
type Strategy string

const (
	win  Strategy = "Z"
	draw Strategy = "Y"
	lose Strategy = "X"
)

func choose(strategy Strategy, opponentShapeName string) string {
	switch strategy {
	case win:
		for idx, s := range shapeOrderedSet {
			if s == opponentShapeName {
				if idx == 0 {
					return shapeOrderedSet[2]
				}
				return shapeOrderedSet[idx-1]
			}
		}
	case lose:
		for idx, s := range shapeOrderedSet {
			if s == opponentShapeName {
				if idx == 2 {
					return shapeOrderedSet[0]
				}
				return shapeOrderedSet[idx+1]
			}
		}
	}
	return opponentShapeName
}

func PartOne(input []string) string {

	var myTotalScore int
	for _, line := range input {

		splitted := strings.Split(line, " ")

		opponentSymbol := splitted[0]
		mySymbol := splitted[1]

		opponentShape := opponentDict[opponentSymbol]
		myShape := myDict[mySymbol]

		myTotalScore += myShape.score
		if myShape.score == opponentShape.score {
			myTotalScore += 3
		}

		if won(myShape.name, opponentShape.name) {
			myTotalScore += 6
		}
	}
	return strconv.Itoa(myTotalScore)
}

func PartTwo(input []string) string {

	var myTotalScore int
	for _, line := range input {

		splitted := strings.Split(line, " ")

		opponentSymbol := splitted[0]
		mySymbol := splitted[1]

		opponentShape := opponentDict[opponentSymbol]

		chosenShapeName := choose(Strategy(mySymbol), opponentShape.name)
		myChosenSymbol := reverseLookUp(chosenShapeName)

		myTotalScore += myDict[myChosenSymbol].score
		if won(chosenShapeName, opponentShape.name) {
			myTotalScore += 6
		}
		if chosenShapeName == opponentShape.name {
			myTotalScore += 3
		}
	}
	return strconv.Itoa(myTotalScore)
}
