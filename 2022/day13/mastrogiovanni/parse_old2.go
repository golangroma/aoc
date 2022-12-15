package main

import "fmt"

func MakeArray(token []Token) []Token {
	return []Token{
		{Type: Open},
		token[0],
		{Type: Close}}
}

func CompareRecursive(left, right [][]Token) int {
	fmt.Printf("A: %+v\n", left)
	fmt.Printf("B: %+v\n", right)
	if len(left) > len(right) {
		fmt.Printf("left size %d > right size %d\n", len(left), len(right))
		return +1
	}
	if len(left) == 0 {
		fmt.Println("Compare Y")
		if len(left) == len(right) {
			return 0
		}
		return -1
	}
	for i := 0; i < len(left); i++ {
		fmt.Printf("%d/%d\n", i, len(left))
		fmt.Printf("Compare A: %+v (%d)\n", left[i], len(left[i]))
		fmt.Printf("Compare B: %+v (%d)\n", right[i], len(right[i]))
		if IsArray(left[i]) {
			if IsArray(right[i]) {
				l := Parse(left[i])
				r := Parse(right[i])
				fmt.Println("Compare 1", l, r)
				res := CompareRecursive(l, r)
				if res != 0 {
					return res
				}
			} else {
				l := Parse(left[i])
				r := Parse(MakeArray(right[i]))
				fmt.Println("Compare 2", l, r)
				res := CompareRecursive(l, r)
				if res != 0 {
					return res
				}
			}
		} else {
			if IsArray(right[i]) {
				l := Parse(MakeArray(left[i]))
				r := Parse(right[i])
				fmt.Println("Compare 3", l, r)
				res := CompareRecursive(l, r)
				if res != 0 {
					return res
				}
			} else {
				l := left[i][0].Value
				r := right[i][0].Value
				fmt.Println("Compare 4", l, r)
				if l < r {
					return -1
				} else if l > r {
					return 1
				}
			}
		}
	}

	return -1
}

func Compare2(left, right string) bool {
	tokensLeft := Parse(Reverse(GetTokens(left)))
	tokensRight := Parse(Reverse(GetTokens(right)))
	return CompareRecursive(tokensLeft, tokensRight) < 0
}
