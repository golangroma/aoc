package main

import (
	"bytes"
	"strconv"
)

func PartOne(input []string) string {
	//
	// DLzSMtDLtzmmlDlMlMDbcrcTDqFvVvVqqTbD
	// Deve essere divisa in
	// DLzSMtDLtzmmlDlMlM
	// DbcrcTDqFvVvVqqTbD
	//

	totalPoints := 0

	for _, v := range input {
		compartments := splitSubN(v, len(v)/2)
		totalPoints += findCommonCharPointsByTwo(compartments[0], compartments[1])
	}

	return strconv.Itoa(totalPoints)
}

func PartTwo(input []string) string {
	//
	// PPZTzDhJPLqPhqDTqrwQZZWbmCBMJMcsNmCBFWmMcsNb
	// vplSlfdfGvfRRGsgNcMglsFWMWMC
	// jtjvFHdjjwqrwqwL
	//

	compartments := [3]string{}
	totalPoints := 0
	elfNumberInGroup := 0

	for k, v := range input {
		compartments[elfNumberInGroup] = v

		if (k+1)%3 == 0 {
			totalPoints += findCommonCharPointsByThree(compartments[0], compartments[1], compartments[2])
			elfNumberInGroup = 0
		} else {
			elfNumberInGroup++
		}
	}

	return strconv.Itoa(totalPoints)
}

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func findCommonCharPointsByTwo(s1, s2 string) int {
	//
	// DLzSMtDLtzmmlDlMlM
	// DbcrcTDqFvVvVqqTbD
	//
	rc := 0
	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))
	for _, r := range runes1 {
		for _, rr := range runes2 {
			if r == rr {
				if rr >= 65 && rr <= 90 {
					// Uppercase
					rc = int(rr) - 38
				} else {
					// Lowercase
					rc = int(rr) - 96
				}
			}
		}
	}

	return rc
}

func findCommonCharPointsByThree(s1, s2, s3 string) int {
	//
	// PPZTzDhJPLqPhqDTqrwQZZWbmCBMJMcsNmCBFWmMcsNb
	// vplSlfdfGvfRRGsgNcMglsFWMWMC
	// jtjvFHdjjwqrwqwL
	//
	rc := 0
	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))
	runes3 := bytes.Runes([]byte(s3))
	for _, r := range runes1 {
		for _, rr := range runes2 {
			for _, rrr := range runes3 {
				if r == rr && rr == rrr {
					if rrr >= 'A' && rrr <= 'Z' {
						// Uppercase
						rc = int(rrr) - 38
					} else {
						// Lowercase
						rc = int(rrr) - 96
					}
				}
			}
		}
	}

	return rc
}
