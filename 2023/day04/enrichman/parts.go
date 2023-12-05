package main

import (
	"log"
	"strconv"
	"strings"
)

type ScratchCard struct {
	ID             int
	WinningNumbers []int
}

func (s ScratchCard) HasNumber(num int) bool {
	for _, wn := range s.WinningNumbers {
		if wn == num {
			return true
		}
	}
	return false
}

func PartOne(input []string) string {
	sum := 0

	for _, line := range input {
		scratch, nums := mustParseLine(line)

		pow := 0

		for _, num := range nums {
			if scratch.HasNumber(num) {
				if pow == 0 {
					pow = 1
				} else {
					pow = pow * 2
				}
			}
		}

		sum += pow
	}

	return strconv.Itoa(sum)
}

func PartTwo(input []string) string {
	copies := make([]int, len(input))

	for i, line := range input {
		scratch, nums := mustParseLine(line)

		counter := 0
		for _, num := range nums {
			if scratch.HasNumber(num) {
				counter++
			}
		}

		// add original
		copies[i] += 1

		// increment copies of winning cards
		start := i + 1
		end := start + counter
		for j := start; j < end; j++ {
			copies[j] += copies[i]
		}
	}

	sum := 0
	for _, c := range copies {
		sum += c
	}
	return strconv.Itoa(sum)
}

func mustParseLine(input string) (ScratchCard, []int) {
	splitted := strings.Split(input, "|")

	scratchCard := mustParseScratchLine(splitted[0])
	numbers := mustParseNumbersLine(splitted[1])

	return scratchCard, numbers
}

func mustParseScratchLine(input string) ScratchCard {
	input = strings.TrimSpace(input)
	splitted := strings.Split(input, ":")

	idStr := strings.Fields(splitted[0])[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	winningNumbers := mustParseNumbersLine(splitted[1])

	return ScratchCard{
		ID:             id,
		WinningNumbers: winningNumbers,
	}
}

func mustParseNumbersLine(input string) []int {
	input = strings.TrimSpace(input)
	numbersArr := strings.Fields(input)

	numbers := []int{}
	for _, numStr := range numbersArr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	return numbers
}
