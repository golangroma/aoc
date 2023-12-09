package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var strength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

const (
	HighCard = iota
	OnePair
	TwoPair
	Three
	FullHouse
	FourOfaKind
	FiveOfaKind
)

func handType(i string) int {
	cards := map[string]int{}
	keys := []string{}
	for _, v := range i {
		cards[string(v)]++
		keys = append(keys, string(v))
	}
	// five of a kind
	if len(cards) == 1 {
		return FiveOfaKind
	}
	// four of a kind or full house
	if len(cards) == 2 {
		v1, v2 := cards[keys[0]], cards[keys[1]]
		if v1 == 4 || v2 == 4 {
			return FourOfaKind
		}
		return FullHouse
	}
	// three of a kind or two of pair
	if len(cards) == 3 {
		v1, v2, v3 := cards[keys[0]], cards[keys[1]], cards[keys[2]]
		if v1 == 3 || v2 == 3 || v3 == 3 {
			return Three
		}
		return TwoPair
	}
	// One pair
	if len(cards) == 4 {
		return OnePair
	}
	return HighCard
}

func handType2(i string) int {
	old := handType(i)
	if old == FiveOfaKind {
		return old
	}
	cards := map[string]int{}
	keys := []string{}
	for _, v := range i {
		cards[string(v)]++
		keys = append(keys, string(v))
	}
	if slices.Contains(keys, "J") {
		switch old {
		case FourOfaKind, FullHouse:
			return FiveOfaKind
		case Three:
			return FourOfaKind
		case TwoPair:
			if cards["J"] == 2 {
				return FourOfaKind
			}
			return FullHouse
		case OnePair:
			return Three
		case HighCard:
			return OnePair
		}

	}
	return old
}

type hand struct {
	value    string
	prize    int
	handType int
}

func PartOne(input []string) string {
	rank := []hand{}
	for _, v := range input {
		hValue, prize, _ := strings.Cut(v, " ")
		rank = append(rank, hand{
			value:    hValue,
			prize:    mustConvertToInt(prize),
			handType: handType(hValue),
		})
	}

	slices.SortFunc(rank, func(a, b hand) int {
		if a.handType > b.handType {
			return 1
		}
		if a.handType < b.handType {
			return -1
		}
		for i := 0; i < len(a.value); i++ {
			sa := strength[string(a.value[i])]
			sb := strength[string(b.value[i])]
			if sa > sb {
				return 1
			}
			if sa < sb {
				return -1
			}
		}
		return 0
	})
	sum := 0
	for i, v := range rank {
		sum += v.prize * (i + 1)
	}

	return fmt.Sprintf("%d", sum)
}

func PartTwo(input []string) string {
	rank := []hand{}
	for _, v := range input {
		hValue, prize, _ := strings.Cut(v, " ")
		rank = append(rank, hand{
			value:    hValue,
			prize:    mustConvertToInt(prize),
			handType: handType2(hValue),
		})
	}

	slices.SortFunc(rank, func(a, b hand) int {
		if a.handType > b.handType {
			return 1
		}
		if a.handType < b.handType {
			return -1
		}
		for i := 0; i < len(a.value); i++ {
			sa := strength[string(a.value[i])]
			sb := strength[string(b.value[i])]
			if sa == strength["J"] {
				sa = 0
			}
			if sb == strength["J"] {
				sb = 0
			}
			if sa > sb {
				return 1
			}
			if sa < sb {
				return -1
			}
		}
		return 0
	})
	sum := 0
	for i, v := range rank {
		sum += v.prize * (i + 1)
	}

	return fmt.Sprintf("%d", sum)
}

func mustConvertToInt(v string) int {
	r, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return r
}
