package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne(inputs []string) string {
	indexCycle := 0
	samplerCycles := []int{20, 60, 100, 140, 180, 220}
	currentCycle := 1
	cycle := samplerCycles[indexCycle]
	tot := 0
	X := 1
	for _, line := range inputs {
		if line == "noop" {
			if currentCycle+1 == cycle {
				tot += X * (currentCycle + 1)
				if indexCycle+1 >= len(samplerCycles) {
					break
				}
				cycle = samplerCycles[indexCycle+1]
				indexCycle++
			}
			currentCycle++
		} else {
			comps := strings.Split(line, " ")
			value, _ := strconv.Atoi(comps[1])
			if currentCycle+2 >= cycle {
				if currentCycle+2 == cycle {
					tot += (X + value) * (currentCycle + 2)
				} else {
					tot += X * (currentCycle + 1)
				}
				if indexCycle+1 >= len(samplerCycles) {
					break
				}
				cycle = samplerCycles[indexCycle+1]
				indexCycle++
			}
			X += value
			currentCycle += 2
		}
		if indexCycle >= len(samplerCycles) {
			break
		}
	}
	return fmt.Sprintf("%d", tot)
}

func PartTwo(inputs []string) string {
	X := 1
	clock := 0
	crt := ""
	value := 0
	length := 1
	for _, line := range inputs {
		// Fetch
		if line == "noop" {
			value = 0
			length = 1
		} else {
			comps := strings.Split(line, " ")
			value, _ = strconv.Atoi(comps[1])
			length = 2
		}

		// Crt draw
		for i := 0; i < length; i++ {
			cursor := clock % 40
			if cursor >= X-1 && cursor <= X+1 {
				crt += "#"
			} else {
				crt += "."
			}
			if i == length-1 {
				X = X + value
			}
			clock++
			if clock%40 < (clock-1)%40 {
				crt += "\n"
			}
		}

	}
	return crt
}
