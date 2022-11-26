package main

import (
	"math"
	"strconv"
)

func PartOne(input []string) string {
	occurrences := findOccurrences(input)

	var gammaRate, epsilonRate string
	for i := range occurrences {
		if occurrences[i] > 0 {
			gammaRate += "0"
			epsilonRate += "1"
		} else {
			gammaRate += "1"
			epsilonRate += "0"
		}
	}

	return strconv.Itoa(binToDec(gammaRate) * binToDec(epsilonRate))
}

func findOccurrences(input []string) []int {
	occurrences := make([]int, len(input[0]))

	for _, binary := range input {
		for i := range binary {
			if string(binary[i]) == "1" {
				occurrences[i] = occurrences[i] + 1
			} else {
				occurrences[i] = occurrences[i] - 1
			}
		}
	}

	return occurrences
}

func binToDec(bin string) int {
	dec := 0

	count := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if string(bin[i]) == "1" {
			dec += int(math.Pow(2, float64(count)))
		}
		count++
	}

	return dec
}

func PartTwo(input []string) string {
	originalInput := input

	// oxigen
	occurrences := findOccurrences(input)
	for i := range occurrences {
		// more "1" than "0"
		if occurrences[i] >= 0 {
			input = filter(input, i, "1")
		} else {
			input = filter(input, i, "0")
		}

		occurrences = findOccurrences(input)
	}
	oxigenRating := input[0]

	// CO2
	input = originalInput
	occurrences = findOccurrences(input)

	for i := range occurrences {
		// more "1" than "0"
		if occurrences[i] < 0 {
			input = filter(input, i, "1")
		} else {
			input = filter(input, i, "0")
		}

		if len(input) == 1 {
			break
		}

		occurrences = findOccurrences(input)
	}
	co2Rating := input[0]

	return strconv.Itoa(binToDec(oxigenRating) * binToDec(co2Rating))
}

func filter(input []string, indexToCheck int, prefix string) []string {
	filtered := []string{}

	for _, line := range input {
		if string(line[indexToCheck]) == prefix {
			filtered = append(filtered, line)
		}
	}

	return filtered
}
