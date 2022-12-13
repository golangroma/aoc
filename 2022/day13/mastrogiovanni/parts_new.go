package main

import "encoding/json"

func parse(line string) []any {
	ll := []any{}
	_ = json.Unmarshal([]byte(line), &ll)
	return ll
}

func intCompare(left, right int) int {
	if left < right {
		return -1
	}
	if left > right {
		return 1
	}
	return 0
}

func min(left, right int) int {
	if left <= right {
		return left
	}
	return right
}

type Items []any

func (a Items) Len() int           { return len(a) }
func (a Items) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Items) Less(i, j int) bool { return compare(a[i], a[j]) < 0 }

func compare(left any, right any) int {
	lInt, lList, lIsInt := typeOf(left)
	rInt, rList, rIsInt := typeOf(right)

	// both integer
	if lIsInt && rIsInt {
		return intCompare(lInt, rInt)
	}

	// both lists
	if !lIsInt && !rIsInt {
		lLen, rLen := len(lList), len(rList)
		for i := 0; i < min(lLen, rLen); i++ {
			// return if they're different
			if cmp := compare(lList[i], rList[i]); cmp != 0 {
				return cmp
			}
			// continue if they're equal
		}

		// when a list is a prefix of another
		// the shortest comes first
		return intCompare(lLen, rLen)
	}

	if lIsInt {
		return compare([]any{lInt}, right)
	}

	if rIsInt {
		return compare(left, []any{rInt})
	}

	panic("wrong types")
}

func typeOf(x any) (int, []any, bool) {
	switch val := x.(type) {
	case int:
		return val, nil, true
	case float64:
		return int(val), nil, true
	case []any:
		return 0, val, false
	}
	panic("wrong type")
}

func Compare3(left, right string) bool {
	tokensLeft := parse(left)
	tokensRight := parse(right)
	return compare(tokensLeft, tokensRight) < 0
}
