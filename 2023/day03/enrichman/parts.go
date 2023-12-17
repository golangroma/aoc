package main

import (
	"log"
	"strconv"
	"unicode"
)

type PartNumber struct {
	Number int
	Line   int
	Start  int
	End    int
}

type Symbol struct {
	Line int
	X    int
	Gear bool
}

func PartOne(input []string) string {
	allParts := []PartNumber{}
	allSymbols := []Symbol{}

	for i, line := range input {
		parts, symbols := mustParseLine(i, line)

		allParts = append(allParts, parts...)
		allSymbols = append(allSymbols, symbols...)
	}

	sum := 0

	// check parts validity
	for _, part := range allParts {
		if part.IsValid(allSymbols) {
			sum += part.Number
		}
	}

	return strconv.Itoa(sum)
}

func PartTwo(input []string) string {
	allParts := []PartNumber{}
	gearsSymbols := []Symbol{}

	for i, line := range input {
		parts, symbols := mustParseLine(i, line)

		allParts = append(allParts, parts...)

		for _, sym := range symbols {
			if sym.Gear {
				gearsSymbols = append(gearsSymbols, sym)
			}
		}
	}

	sum := 0

	// find near parts
	for _, gear := range gearsSymbols {
		sum += gear.Ratio(allParts)
	}

	return strconv.Itoa(sum)
}

// returns the PartNumbers and Symbols for that line
func mustParseLine(lineNumber int, input string) ([]PartNumber, []Symbol) {
	parts := []PartNumber{}
	symbols := []Symbol{}

	var partNumber *PartNumber
	var numStr string

	for i, s := range input {
		// part number
		if unicode.IsDigit(s) {
			// initialize part if nil
			if partNumber == nil {
				partNumber = &PartNumber{Line: lineNumber, Start: i}
			}
			numStr += string(s)
			continue
		}

		// not part number, finalize if any
		if numStr != "" {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}

			partNumber.Number = num
			partNumber.End = i - 1
			parts = append(parts, *partNumber)

			partNumber = nil
			numStr = ""
		}

		if s != '.' {
			sym := Symbol{
				X:    i,
				Line: lineNumber,
				Gear: (s == '*'),
			}
			symbols = append(symbols, sym)
		}
	}

	//  finalize if any
	if numStr != "" {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}

		partNumber.Number = num
		partNumber.End = len(input) - 1
		parts = append(parts, *partNumber)
	}

	return parts, symbols
}

// check if there is a symbol near the part number
func (p PartNumber) IsValid(symbols []Symbol) bool {
	for _, s := range symbols {
		// same line, check if near left or right
		if s.Line == p.Line {
			// symbol on left
			if s.X == p.Start-1 {
				return true
			}
			// symbol on right
			if s.X == p.End+1 {
				return true
			}
		}

		// symbol on line above or below
		if s.Line-1 == p.Line || s.Line+1 == p.Line {
			// symbol has to be on start-1 and end+1
			if s.X >= p.Start-1 && s.X <= p.End+1 {
				return true
			}
		}
	}

	return false
}

// check if there is a symbol near the part number
func (s Symbol) Ratio(parts []PartNumber) int {
	nearParts := []PartNumber{}

	for _, p := range parts {
		// same line, check if near left or right
		if s.Line == p.Line {
			// symbol on left
			if s.X == p.Start-1 {
				nearParts = append(nearParts, p)
			}
			// symbol on right
			if s.X == p.End+1 {
				nearParts = append(nearParts, p)
			}
		}

		// symbol on line above or below
		if s.Line-1 == p.Line || s.Line+1 == p.Line {
			// symbol has to be on start-1 and end+1
			if s.X >= p.Start-1 && s.X <= p.End+1 {
				nearParts = append(nearParts, p)
			}
		}
	}

	if len(nearParts) > 2 {
		log.Fatal("too many near parts")
	}

	if len(nearParts) == 2 {
		return nearParts[0].Number * nearParts[1].Number
	}

	return 0
}
