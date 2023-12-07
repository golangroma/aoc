package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func PartOne(input []string) string {
	var sum int
	for _, v := range input {
		_, nextToken, _ := strings.Cut(v, ": ")
		winningBoard, myNumbers, _ := strings.Cut(nextToken, " | ")
		wBoard, gNumbers := parseBoard(winningBoard), parseBoard(myNumbers)
		winningNumbers, points := []int{}, 0
		for _, n := range gNumbers {
			if slices.Contains(wBoard, n) {
				winningNumbers = append(winningNumbers, n)
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		//fmt.Println(points)
		sum += points
	}

	return fmt.Sprintf("%v", sum)
}

type Card struct {
	id             int
	winningNumbers []int
	guessedNumbers []int
}

func (c Card) Play() int {
	sum := 0
	for _, n := range c.winningNumbers {
		if slices.Contains(c.winningNumbers, n) {
			sum++
		}
	}
	return sum
}

func PartTwo(input []string) string {
	var sum int
	cardCopies := map[int]int{}
	parsedCards := map[int]Card{}
	latestCardID := 0

	for _, v := range input {
		var gameID int
		game, nextToken, _ := strings.Cut(v, ": ")
		_, _ = fmt.Sscanf(game, "Card %d", &gameID)
		latestCardID = gameID
		winningBoard, myNumbers, _ := strings.Cut(nextToken, " | ")
		wBoard, gNumbers := parseBoard(winningBoard), parseBoard(myNumbers)
		parsedCards[gameID] = Card{
			id:             latestCardID,
			winningNumbers: wBoard,
			guessedNumbers: gNumbers,
		}
	}

	for _, v := range parsedCards {
		n := v.Play()
		for i := 0; i <= n; i++ {
			if v.id+i > latestCardID {
				break
			}
			cardCopies[v.id+i] = cardCopies[v.id+i] + 1
		}
	}

	for k, v := range cardCopies {
		if k > latestCardID {
			continue
		}
		sum += v + 1
	}

	return fmt.Sprintf("%v", sum)
}

func parseBoard(s string) []int {
	result := []int{}
	for _, v := range strings.Split(s, " ") {
		if v != "" {
			c, _ := strconv.Atoi(v)
			result = append(result, c)
		}
	}
	return result
}
