package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Garden struct {
	seeds                 []int
	seedToSoil            []mapInfo
	soilToFeritilizer     []mapInfo
	fertilizerToWater     []mapInfo
	waterToLight          []mapInfo
	lightToTemperature    []mapInfo
	temperatureToHumidity []mapInfo
	humidityToLocation    []mapInfo
}

type mapInfo struct {
	destination, source, length int
}

func PartOne(input []string) string {
	g := Parse(input)
	seedsLocation := map[int]int{}
	mapQueue := [][]mapInfo{
		g.seedToSoil,
		g.soilToFeritilizer,
		g.fertilizerToWater,
		g.waterToLight,
		g.lightToTemperature,
		g.temperatureToHumidity,
		g.humidityToLocation,
	}
	for _, seed := range g.seeds {
		location := -1
		target := seed
		for j, mapFromQueue := range mapQueue {
			found := false
			candidateValue := 0
			for i := 0; i < len(mapFromQueue); i++ {
				ok, n := isInRange(target, mapFromQueue[i])
				if ok {
					candidateValue = n
					target = n
					found = true
					break
				}
			}
			if !found {
				candidateValue = target
			}
			if j == len(mapQueue)-1 {
				location = candidateValue
			}
		}
		seedsLocation[seed] = location
	}

	locations := []int{}
	for _, location := range seedsLocation {
		locations = append(locations, location)
	}
	slices.Sort(locations)
	return fmt.Sprintf("%v", locations[0])
}

func isInRange(target int, m mapInfo) (bool, int) {
	offset := target - m.source
	if target >= m.source && target < m.source+m.length {
		return true, m.destination + offset
	}
	return false, -1
}

func Parse(input []string) Garden {
	g := Garden{}
	seeds := strings.Split(input[0], " ")
	for _, v := range seeds[1:] {
		g.seeds = append(g.seeds, mustConvertToInt(v))
	}
	mapType := ""
	for _, v := range input[2:] {
		if v == "" {
			continue
		}
		if strings.HasSuffix(v, "map:") {
			mapType = v
			continue
		}
		switch mapType {
		case "seed-to-soil map:":
			g.seedToSoil = append(g.seedToSoil, parseMap(v))
		case "soil-to-fertilizer map:":
			g.soilToFeritilizer = append(g.soilToFeritilizer, parseMap(v))
		case "fertilizer-to-water map:":
			g.fertilizerToWater = append(g.fertilizerToWater, parseMap(v))
		case "water-to-light map:":
			g.waterToLight = append(g.waterToLight, parseMap(v))
		case "light-to-temperature map:":
			g.lightToTemperature = append(g.lightToTemperature, parseMap(v))
		case "temperature-to-humidity map:":
			g.temperatureToHumidity = append(g.temperatureToHumidity, parseMap(v))
		case "humidity-to-location map:":
			g.humidityToLocation = append(g.humidityToLocation, parseMap(v))
		}
	}
	return g
}

func parseMap(m string) mapInfo {
	i := mapInfo{}
	fmt.Sscanf(m, "%d %d %d", &i.destination, &i.source, &i.length)
	return i
}

func PartTwo(input []string) string {
	g := Parse(input)
	mapQueue := [][]mapInfo{
		g.seedToSoil,
		g.soilToFeritilizer,
		g.fertilizerToWater,
		g.waterToLight,
		g.lightToTemperature,
		g.temperatureToHumidity,
		g.humidityToLocation,
	}
	// I need to go to lunch let's brute force IDK it will take more than 1 hour :)
	minLocation := math.MaxInt
	for i := 0; i < len(g.seeds)-1; i += 2 {
		s, e := g.seeds[i], g.seeds[i+1]
		for e >= 0 {
			seed := s + e
			location := -1
			target := seed
			for j, mapFromQueue := range mapQueue {
				found := false
				candidateValue := 0
				for i := 0; i < len(mapFromQueue); i++ {
					ok, n := isInRange(target, mapFromQueue[i])
					if ok {
						candidateValue = n
						target = n
						found = true
						break
					}
				}
				if !found {
					candidateValue = target
				}
				if j == len(mapQueue)-1 {
					location = candidateValue
				}
			}
			if location < minLocation {
				minLocation = location
			}
			e--
		}
	}
	return fmt.Sprintf("%v", minLocation)
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
