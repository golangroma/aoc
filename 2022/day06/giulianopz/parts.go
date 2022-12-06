package main

import "fmt"

var occurences = make(map[byte]int, 0)

func PartOne(input []string) string {

	maxChars := 3
	var processedCharsNum int
	signals := input[0]
	for i := 0; i < len(signals); i++ {

		if i >= maxChars {

			incrByOne(signals[i], signals[i-1], signals[i-2], signals[i-3])

			if allOnce(signals[i], signals[i-1], signals[i-2], signals[i-3]) {
				processedCharsNum = i + 1
				break
			}
		}
	}
	return fmt.Sprintf("%d", processedCharsNum)
}

func PartTwo(input []string) string {

	maxChars := 13
	var processedCharsNum int
	signals := input[0]
	for i := 0; i < len(signals); i++ {

		if i >= maxChars {

			lowerBound := i - maxChars
			incrByOne([]byte(signals)[lowerBound:i]...)

			if allOnce([]byte(signals)[lowerBound:i]...) {
				processedCharsNum = i + 1
				break
			}
		}
	}
	return fmt.Sprintf("%d", processedCharsNum)
}

func incrByOne(chars ...byte) {
	for _, c := range chars {
		occurences[c] += 1
	}
}

func allOnce(chars ...byte) bool {
	for _, c := range chars {
		if occurences[c] != 1 {
			occurences = make(map[byte]int, 0)
			return false
		}
	}
	return true
}
