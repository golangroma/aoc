package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func prepend(element string, data []string) []string {
	return append([]string{element}, data...)
}

func PartOne(input []string) string {

	re, err := regexp.Compile(`\w+\s(\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	crates := make(map[int][]string, 0)
	row := 0
	for _, line := range input {

		if strings.Index(line, "[") != -1 {
			for i := 0; i < len(line); i += 4 {

				if string(line[i]) == "[" {
					letter := string(line[i+1])

					if letter != "" {
						elements, ok := crates[row]
						if !ok {
							crates[row] = make([]string, 0)
						}
						crates[row] = prepend(letter, elements)
					}
				}
				row++
			}
			row = 0
		}

		if re.MatchString(line) {
			matches := re.FindAllStringSubmatch(line, -1)

			num := asInt(matches[0][1])
			from := asInt(matches[1][1]) - 1
			to := asInt(matches[2][1]) - 1

			source := crates[from]

			moving := source[len(source)-num:]
			for len(moving) != 0 {
				popped := moving[len(moving)-1]
				crates[to] = append(crates[to], popped)
				moving = moving[:len(moving)-1]
			}
			crates[from] = source[:len(source)-num]
		}
	}

	var top string
	for i := 0; i < len(crates); i++ {
		top += crates[i][len(crates[i])-1]
	}
	return top
}

func PartTwo(input []string) string {

	re, err := regexp.Compile(`\w+\s(\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	crates := make(map[int][]string, 0)
	row := 0
	for _, line := range input {

		if strings.Index(line, "[") != -1 {
			for i := 0; i < len(line); i += 4 {

				if string(line[i]) == "[" {
					letter := string(line[i+1])

					if letter != "" {
						elements, ok := crates[row]
						if !ok {
							crates[row] = make([]string, 0)
						}
						crates[row] = prepend(letter, elements)
					}
				}
				row++
			}
			row = 0
		}

		if re.MatchString(line) {
			matches := re.FindAllStringSubmatch(line, -1)

			num := asInt(matches[0][1])
			from := asInt(matches[1][1]) - 1
			to := asInt(matches[2][1]) - 1

			source := crates[from]

			moving := source[len(source)-num:]
			crates[to] = append(crates[to], moving...)
			crates[from] = source[:len(source)-num]
		}
	}

	var top string
	for i := 0; i < len(crates); i++ {
		top += crates[i][len(crates[i])-1]
	}
	return top
}

func asInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
