package main

import (
	"fmt"
	"strconv"
)

type Schematic struct {
	value        string
	isAdjacent   bool
	gearPosition Coord
}

type Coord struct {
	x, y int
}

func PartOne(input []string) string {
	var sum int
	schematics := []Schematic{}
	for i, line := range input {
		schematicNew := Schematic{}
		add := func() {
			if schematicNew.value != "" {
				schematics = append(schematics, schematicNew)
				if schematicNew.isAdjacent {
					v, _ := strconv.Atoi(schematicNew.value)
					sum += v
				}
				schematicNew = Schematic{}
			}
		}
		for j, char := range line {
			if char >= '0' && char <= '9' {
				schematicNew.value += string(char)

				topIdx := i - 1
				bottomIdx := i + 1
				leftIdx := j - 1
				rightIdx := j + 1

				// top
				if topIdx >= 0 {
					if isSpecialChar([]rune(input[topIdx])[j]) {
						schematicNew.isAdjacent = true
					}
				}
				// bottom
				if bottomIdx < len(input) {
					if isSpecialChar([]rune(input[bottomIdx])[j]) {
						schematicNew.isAdjacent = true
					}
				}
				// left
				if leftIdx >= 0 {
					if isSpecialChar([]rune(line)[leftIdx]) {
						schematicNew.isAdjacent = true
					}
				}
				// right
				if rightIdx < len(line) {
					if isSpecialChar([]rune(line)[rightIdx]) {
						schematicNew.isAdjacent = true
					}
				}
				// top-left
				if topIdx >= 0 && leftIdx >= 0 {
					if isSpecialChar([]rune(input[topIdx])[leftIdx]) {
						schematicNew.isAdjacent = true
					}
				}
				// top-right
				if topIdx >= 0 && rightIdx < len(input[topIdx]) {
					if isSpecialChar([]rune(input[topIdx])[rightIdx]) {
						schematicNew.isAdjacent = true
					}
				}
				// bottom-left
				if bottomIdx < len(input) && leftIdx >= 0 {
					if isSpecialChar([]rune(input[bottomIdx])[leftIdx]) {
						schematicNew.isAdjacent = true
					}
				}
				// bottom-right
				if bottomIdx < len(input) && rightIdx < len(input[bottomIdx]) {
					if isSpecialChar([]rune(input[bottomIdx])[rightIdx]) {
						schematicNew.isAdjacent = true
					}
				}
			} else {
				add()
			}
		}
		add()
	}

	return fmt.Sprintf("%v", sum)
}

func isSpecialChar(v rune) bool {
	return (v < '0' || v > '9') && v != '.'
}

func PartTwo(input []string) string {
	var sum int
	schematics := []Schematic{}
	sharedGears := map[Coord][]Schematic{}
	for i, line := range input {
		schematicNew := Schematic{}
		add := func() {
			if schematicNew.value != "" {
				if schematicNew.isAdjacent {
					schematics = append(schematics, schematicNew)
					sharedGears[schematicNew.gearPosition] = append(sharedGears[schematicNew.gearPosition], schematicNew)
				}
				schematicNew = Schematic{}
			}
		}
		for j, char := range line {
			if char >= '0' && char <= '9' {
				schematicNew.value += string(char)

				topIdx := i - 1
				bottomIdx := i + 1
				leftIdx := j - 1
				rightIdx := j + 1

				// top
				if topIdx >= 0 {
					if isGear([]rune(input[topIdx])[j]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = topIdx
						schematicNew.gearPosition.y = j
					}
				}
				// bottom
				if bottomIdx < len(input) {
					if isGear([]rune(input[bottomIdx])[j]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = bottomIdx
						schematicNew.gearPosition.y = j
					}
				}
				// left
				if leftIdx >= 0 {
					if isGear([]rune(line)[leftIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = i
						schematicNew.gearPosition.y = leftIdx
					}
				}
				// right
				if rightIdx < len(line) {
					if isGear([]rune(line)[rightIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = i
						schematicNew.gearPosition.y = rightIdx
					}
				}
				// top-left
				if topIdx >= 0 && leftIdx >= 0 {
					if isGear([]rune(input[topIdx])[leftIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = topIdx
						schematicNew.gearPosition.y = leftIdx
					}
				}
				// top-right
				if topIdx >= 0 && rightIdx < len(input[topIdx]) {
					if isGear([]rune(input[topIdx])[rightIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = topIdx
						schematicNew.gearPosition.y = rightIdx
					}
				}
				// bottom-left
				if bottomIdx < len(input) && leftIdx >= 0 {
					if isGear([]rune(input[bottomIdx])[leftIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = bottomIdx
						schematicNew.gearPosition.y = leftIdx
					}
				}
				// bottom-right
				if bottomIdx < len(input) && rightIdx < len(input[bottomIdx]) {
					if isGear([]rune(input[bottomIdx])[rightIdx]) {
						schematicNew.isAdjacent = true
						schematicNew.gearPosition.x = bottomIdx
						schematicNew.gearPosition.y = rightIdx
					}
				}
			} else {
				add()
			}
		}
		add()
	}
	for _, v := range sharedGears {
		if len(v) <= 1 || len(v) > 2 {
			continue
		}
		sum += mustConvertToInt(v[0].value) * mustConvertToInt(v[1].value)
	}

	return fmt.Sprintf("%v", sum)
}

func isGear(v rune) bool {
	return v == '*'
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
