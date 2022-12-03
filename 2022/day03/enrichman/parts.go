package main

import (
	"strconv"
	"unicode"
)

func PartOne(input []string) string {
	commonItems := ""

	for _, rucksack := range input {
		compartment1 := make(map[string]struct{})

		for i, item := range rucksack {
			// first compartment
			if i < len(rucksack)/2 {
				compartment1[string(item)] = struct{}{}
			} else { // second compartment
				if _, found := compartment1[string(item)]; found {
					commonItems += string(item)
					break
				}
			}
		}

	}

	return strconv.Itoa(priority(commonItems))
}

func priority(items string) int {
	priority := 0

	for _, r := range items {
		if unicode.IsUpper(r) {
			//fmt.Println('A' - 38)
			priority = priority + int(r) - 38
		} else {
			//fmt.Println('a' - 96)
			priority = priority + int(r) - 96
		}
	}

	return priority
}

type Group struct {
	ID        string
	Rucksacks []string
}

func NewGroup(rucksacks ...string) *Group {
	ids := uniqueMap(rucksacks[0])

	for i := 1; i < len(rucksacks); i++ {
		ids = commonItems(
			ids,
			uniqueMap(rucksacks[i]),
		)
	}

	var id string
	for k := range ids {
		id += k
	}

	return &Group{
		ID:        id,
		Rucksacks: rucksacks,
	}
}

func uniqueMap(items string) map[string]struct{} {
	unique := make(map[string]struct{})
	for _, item := range items {
		unique[string(item)] = struct{}{}
	}
	return unique
}

func commonItems(m1, m2 map[string]struct{}) map[string]struct{} {
	commonItems := make(map[string]struct{})
	for k := range m1 {
		if _, found := m2[k]; found {
			commonItems[k] = struct{}{}
		}
	}
	return commonItems
}

func PartTwo(input []string) string {
	groups := []*Group{}

	var group *Group
	for i := 0; i < len(input); i += 3 {
		if i%3 == 0 {
			group = NewGroup(input[i], input[i+1], input[i+2])
			groups = append(groups, group)
		}
	}

	totPriority := 0
	for _, g := range groups {
		totPriority += priority(g.ID)
	}

	return strconv.Itoa(totPriority)
}
