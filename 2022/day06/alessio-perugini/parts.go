package main

import "fmt"

func PartOne(input []string) string {
	return fmt.Sprintf("%d", findFirstUniqueSequenceOfXChars(input[0], 4))
}

func PartTwo(input []string) string {
	return fmt.Sprintf("%d", findFirstUniqueSequenceOfXChars(input[0], 14))
}

func findFirstUniqueSequenceOfXChars(input string, x int) int {
	occurrences := make(map[byte]int)
	for i := 0; i < x; i++ {
		occurrences[input[i]]++
	}

	for i := x; i < len(input); i++ {
		if len(occurrences) == x && isUnique(occurrences) {
			return i
		}
		firstElementOfSequence := input[i-x]
		if v, ok := occurrences[firstElementOfSequence]; ok {
			if v > 1 {
				occurrences[firstElementOfSequence]--
			} else {
				delete(occurrences, firstElementOfSequence)
			}
		}
		occurrences[input[i]]++
	}
	return -1
}

func isUnique(m map[byte]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
