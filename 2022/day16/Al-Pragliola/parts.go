package main

import (
	"strconv"
	"strings"
)

type PressureFinder struct {
	distances map[string]map[string]int
	flowRates map[string]int
	timeLimit int
}

func NewPressureFinder(distances map[string]map[string]int, flowRates map[string]int, timeLimit int) *PressureFinder {
	return &PressureFinder{distances: distances, flowRates: flowRates, timeLimit: timeLimit}
}

func (pf *PressureFinder) FindMaxPressure(
	currentTime int,
	currentPressure int,
	currentFlow int,
	currentValve string,
	remainingValves []string,
) int {
	max := currentPressure + (pf.timeLimit-currentTime)*currentFlow

	for _, v := range remainingValves {
		cost := pf.distances[currentValve][v] + 1

		if cost+currentTime < pf.timeLimit {
			nextCost := pf.FindMaxPressure(
				currentTime+cost,
				currentPressure+currentFlow*cost,
				currentFlow+pf.flowRates[v],
				v,
				RemoveValve(remainingValves, v),
			)

			if nextCost > max {
				max = nextCost
			}
		}
	}

	return max
}

func GetValvesDataAndFlowRates(input []string) (map[string][]*Tile, map[string]int) {
	valvesData := make(map[string][]*Tile)
	flowRates := make(map[string]int)

	for _, line := range input {
		spacesSpl := strings.Split(line, " ")
		name := spacesSpl[1]
		flowRateSpl := strings.Split(spacesSpl[4], "=")[1]
		flowRate, _ := strconv.Atoi(flowRateSpl[:len(flowRateSpl)-1])
		nextValves := spacesSpl[9:]

		valvesData[name] = make([]*Tile, 0)

		for i := 0; i < len(nextValves); i++ {
			t := &Tile{Name: strings.TrimSuffix(nextValves[i], ",")}
			valvesData[name] = append(valvesData[name], t)
		}

		flowRates[name] = flowRate
	}

	return valvesData, flowRates
}

func RemoveValve(valves []string, valve string) []string {
	var result []string

	for i, v := range valves {
		if v == valve {
			result = append(result, valves[:i]...)
			result = append(result, valves[i+1:]...)
			break
		}
	}

	return result
}

func PartOne(input []string) string {
	valves, flowRates := GetValvesDataAndFlowRates(input)

	prunedValves := make([]string, 0)

	for k := range valves {
		if flowRates[k] != 0 {
			prunedValves = append(prunedValves, k)
		}
	}

	distances := make(map[string]map[string]int, len(valves))

	for i := range valves {
		distances[i] = make(map[string]int, len(valves))
	}

	for k1 := range valves {
		for k2 := range valves {
			from := &Tile{Name: k1}
			to := &Tile{Name: k2}

			_, dist, _ := Path(from, to, valves)
			distances[k1][k2] = int(dist)
		}
	}

	pf := NewPressureFinder(distances, flowRates, 30)

	maxPressure := pf.FindMaxPressure(
		0,
		0,
		0,
		"AA",
		prunedValves,
	)

	return strconv.Itoa(maxPressure)
}

func PartTwo(input []string) string {
	valves, flowRates := GetValvesDataAndFlowRates(input)

	prunedValves := make([]string, 0)

	for k := range valves {
		if flowRates[k] != 0 {
			prunedValves = append(prunedValves, k)
		}
	}

	distances := make(map[string]map[string]int, len(valves))

	for i := range valves {
		distances[i] = make(map[string]int, len(valves))
	}

	for k1 := range valves {
		for k2 := range valves {
			from := &Tile{Name: k1}
			to := &Tile{Name: k2}

			_, dist, _ := Path(from, to, valves)
			distances[k1][k2] = int(dist)
		}
	}

	maxPressure := 0

	pf := NewPressureFinder(distances, flowRates, 26)

	for i := 1; i <= len(prunedValves); i++ {
		combos := GetCombos(prunedValves, i)

		for _, c := range combos {
			outer := TakeOuter(c, prunedValves)

			maxPressureMe := pf.FindMaxPressure(
				0,
				0,
				0,
				"AA",
				c,
			)
			maxPressureEl := pf.FindMaxPressure(
				0,
				0,
				0,
				"AA",
				outer,
			)

			if maxPressureMe+maxPressureEl > maxPressure {
				maxPressure = maxPressureMe + maxPressureEl
			}
		}
	}

	return strconv.Itoa(maxPressure)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TakeOuter(s []string, outer []string) []string {
	var result []string

	for _, v := range outer {
		if !Contains(s, v) {
			result = append(result, v)
		}
	}

	return result
}

// GetCombos taken from https://stackoverflow.com/questions/54054062/algorithm-to-generate-k-elements-from-slice-of-n-elements
func GetCombos(set []string, depth int) [][]string {
	return GetCombosHelper(set, depth, 0, []string{}, [][]string{})
}

// GetCombosHelper taken from https://stackoverflow.com/questions/54054062/algorithm-to-generate-k-elements-from-slice-of-n-elements
func GetCombosHelper(set []string, depth int, start int, prefix []string, accum [][]string) [][]string {
	if depth == 0 {
		return append(accum, prefix)
	} else {
		for i := start; i <= len(set)-depth; i++ {
			p := append(prefix, set[i])
			accum = GetCombosHelper(set, depth-1, i+1, p, accum)
		}
		return accum
	}
}
