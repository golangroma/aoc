package main

import (
	"strconv"
)

func PartOne(input []string) string {
	result := len(input)*2 + (len(input[0])-2)*2

	isVisible := make([][]bool, len(input))

	for i := range isVisible {
		isVisible[i] = make([]bool, len(input[0]))
	}

	// L -> R
	for i := 1; i < len(input)-1; i++ {
		localMax := input[i][0]

		for j := 1; j < len(input[0])-1; j++ {
			if input[i][j] > localMax {
				if !isVisible[i][j] {
					result++
					isVisible[i][j] = true
				}

				localMax = input[i][j]
			}
		}
	}

	// R -> L
	for i := 1; i < len(input)-1; i++ {
		localMax := input[i][len(input[0])-1]

		for j := len(input[0]) - 2; j > 0; j-- {
			if input[i][j] > localMax {
				if !isVisible[i][j] {
					result++
					isVisible[i][j] = true
				}

				localMax = input[i][j]
			}
		}
	}

	// T -> B
	for j := 1; j < len(input[0])-1; j++ {
		localMax := input[0][j]

		for i := 1; i < len(input)-1; i++ {
			if input[i][j] > localMax {
				if !isVisible[i][j] {
					result++
					isVisible[i][j] = true
				}

				localMax = input[i][j]
			}
		}
	}

	// B -> T
	for j := 1; j < len(input[0])-1; j++ {
		localMax := input[len(input)-1][j]

		for i := len(input) - 2; i > 0; i-- {
			if input[i][j] > localMax {
				if !isVisible[i][j] {
					result++
					isVisible[i][j] = true
				}

				localMax = input[i][j]
			}
		}
	}

	return strconv.Itoa(result)
}

func PartTwo(input []string) string {
	result := 0

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			s := calcScenery(input, i, j)

			if s > result {
				result = s
			}
		}
	}

	return strconv.Itoa(result)
}

func calcScenery(input []string, i, j int) int {
	var l, r, t, b int

	// L -> R
	for k := j - 1; k >= 0; k-- {
		l++

		if input[i][j] <= input[i][k] {
			break
		}
	}

	// R -> L
	for k := j + 1; k < len(input[0]); k++ {
		r++

		if input[i][j] <= input[i][k] {
			break
		}
	}

	// T -> B
	for k := i - 1; k >= 0; k-- {
		t++

		if input[i][j] <= input[k][j] {
			break
		}
	}

	// B -> T
	for k := i + 1; k < len(input); k++ {
		b++

		if input[i][j] <= input[k][j] {
			break
		}
	}

	return l * r * t * b
}
