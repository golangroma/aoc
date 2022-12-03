package main

import (
	"bytes"
	"strconv"
)

type LowercasePriority int

const (
	unknown LowercasePriority = iota
	a
	b
	c
	d
	e
	f
	g
	h
	i
	j
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
)

type UppercasePriority int

const (
	Unknown UppercasePriority = iota + 26
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

var lowercaseMap = make(map[string]LowercasePriority)
var uppercaseMap = make(map[string]UppercasePriority)

func PartOne(input []string) string {
	//
	// DLzSMtDLtzmmlDlMlMDbcrcTDqFvVvVqqTbD
	// Deve essere divisa in
	// DLzSMtDLtzmmlDlMlM
	// DbcrcTDqFvVvVqqTbD
	//
	initMaps()

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
	initMaps()

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

func initMaps() {
	lowercaseMap["a"] = a
	lowercaseMap["b"] = b
	lowercaseMap["c"] = c
	lowercaseMap["d"] = d
	lowercaseMap["e"] = e
	lowercaseMap["f"] = f
	lowercaseMap["g"] = g
	lowercaseMap["h"] = h
	lowercaseMap["i"] = i
	lowercaseMap["j"] = j
	lowercaseMap["k"] = k
	lowercaseMap["l"] = l
	lowercaseMap["m"] = m
	lowercaseMap["n"] = n
	lowercaseMap["o"] = o
	lowercaseMap["p"] = p
	lowercaseMap["q"] = q
	lowercaseMap["r"] = r
	lowercaseMap["s"] = s
	lowercaseMap["t"] = t
	lowercaseMap["u"] = u
	lowercaseMap["v"] = v
	lowercaseMap["w"] = w
	lowercaseMap["x"] = x
	lowercaseMap["y"] = y
	lowercaseMap["z"] = z

	uppercaseMap["A"] = A
	uppercaseMap["B"] = B
	uppercaseMap["C"] = C
	uppercaseMap["D"] = D
	uppercaseMap["E"] = E
	uppercaseMap["F"] = F
	uppercaseMap["G"] = G
	uppercaseMap["H"] = H
	uppercaseMap["I"] = I
	uppercaseMap["J"] = J
	uppercaseMap["K"] = K
	uppercaseMap["L"] = L
	uppercaseMap["M"] = M
	uppercaseMap["N"] = N
	uppercaseMap["O"] = O
	uppercaseMap["P"] = P
	uppercaseMap["Q"] = Q
	uppercaseMap["R"] = R
	uppercaseMap["S"] = S
	uppercaseMap["T"] = T
	uppercaseMap["U"] = U
	uppercaseMap["V"] = V
	uppercaseMap["W"] = W
	uppercaseMap["X"] = X
	uppercaseMap["Y"] = Y
	uppercaseMap["Z"] = Z
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
					rc = int(uppercaseMap[string(rr)])
				} else {
					// Lowercase
					rc = int(lowercaseMap[string(rr)])
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
					if rrr >= 65 && rrr <= 90 {
						// Uppercase
						rc = int(uppercaseMap[string(rrr)])
					} else {
						// Lowercase
						rc = int(lowercaseMap[string(rrr)])
					}
				}
			}
		}
	}

	return rc
}
