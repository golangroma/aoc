package main

import "fmt"

func AllDifferent(slot string) bool {
	m := make(map[rune]struct{})
	for _, r := range slot {
		_, ok := m[r]
		if ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func FindMarker(msg string, length int) int {
	for i := 0; i < len(msg)-length; i++ {
		if AllDifferent(msg[i : i+length]) {
			return i + length
		}
	}
	return 0
}

func PartOne(input []string) string {
	return fmt.Sprintf("%d", FindMarker(input[0], 4))
}

func PartTwo(input []string) string {
	return fmt.Sprintf("%d", FindMarker(input[0], 14))
}
