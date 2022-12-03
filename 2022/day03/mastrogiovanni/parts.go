package main

import (
	"fmt"
)

func GetCommonChar(line string) rune {
	present := make(map[rune]struct{})
	for i, letter := range line {
		if i < len(line)/2 {
			present[letter] = struct{}{}
		} else {
			_, ok := present[letter]
			if ok {
				return letter
			}
		}
	}
	return rune(' ')
}

func GetCommonChars(lines []string, size int) rune {
	present := make(map[rune]int)
	for _, line := range lines {
		currentGroup := make(map[rune]struct{})
		for _, letter := range line {
			_, ok := currentGroup[letter]
			if !ok {
				currentGroup[letter] = struct{}{}
				value, ok := present[letter]
				if !ok {
					present[letter] = 1
				} else {
					present[letter] = value + 1
				}
			}
		}
	}
	for k, v := range present {
		if v == size {
			return k
		}
	}
	return ' '
}

func Score(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - int('a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c) - int('A') + 27
	}
	return 0
}

func PartOne(input []string) string {
	sum := 0
	for _, row := range input {
		c := GetCommonChar(row)
		sum += Score(c)
	}
	return fmt.Sprintf("%d", sum)
}

func PartTwo(input []string) string {
	size := 3
	sum := 0
	for i := 0; i < len(input); i += size {
		c := GetCommonChars(input[i:i+size], size)
		s := Score(c)
		sum += s
	}
	return fmt.Sprintf("%d", sum)
}
