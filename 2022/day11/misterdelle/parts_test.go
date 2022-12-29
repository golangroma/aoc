package main

import (
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = `Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "10605",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			if got := PartOne(input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "2713310158",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.SplitInput(tc.input)
			if got := PartTwo(input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestApply(t *testing.T) {
	//
	//Starting items: 79, 98
	//Operation: new = old * 19
	//

	tt := []struct {
		name       string
		expression string
		input      int
		expected   int
	}{
		{
			name:       "TestMultiply79*19",
			expression: "old * 19",
			input:      79,
			expected:   500,
		},
		{
			name:       "TestMultiply98*19",
			expression: "old * 19",
			input:      98,
			expected:   620,
		},
		{
			name:       "TestIncrease54By6",
			expression: "old + 6",
			input:      54,
			expected:   20,
		},
		{
			name:       "TestIncrease65By6",
			expression: "old + 6",
			input:      65,
			expected:   23,
		},
		{
			name:       "TestMultiply79ByItself",
			expression: "old * old",
			input:      79,
			expected:   2080,
		},
		{
			name:       "TestIncrease45ByItself",
			expression: "old + old",
			input:      45,
			expected:   30,
		},
		{
			name:       "TestKO",
			expression: "old - old",
			input:      40,
			expected:   0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			op := NewOperation(tc.expression)
			if got := op.applyPartOne(tc.input); got != tc.expected {
				t.Errorf("applyPartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestParseItems(t *testing.T) {
	// Starting items: 79, 98
	got := parseItems("Starting items: 79, 98")
	expected := []int{79, 98}

	if len(got) != len(expected) {
		t.Errorf("parseItems() = %v, want %v", got, expected)
	}

	if got[0] != expected[0] {
		t.Errorf("parseItems() = %v, want %v", got, expected)
	}

	if got[1] != expected[1] {
		t.Errorf("parseItems() = %v, want %v", got, expected)
	}
}

func TestParseOperation(t *testing.T) {
	// Operation: new = old * 19
	got := parseOperation("Operation: new = old * 19")
	expected := "old * 19"

	if got != expected {
		t.Errorf("parseOperation() = %v, want %v", got, expected)
	}
}

func TestParseTest(t *testing.T) {
	// Test: divisible by 23
	got := parseTest("Test: divisible by 23")
	expected := 23

	if got != expected {
		t.Errorf("parseTest() = %v, want %v", got, expected)
	}
}

func TestInspectItems(t *testing.T) {
	monkey0 := NewMonkey(79, 98)
	monkey0.Op = *NewOperation("old * 19")
	monkey0.Divisor = 23
	monkey0.TrueIdx = 2
	monkey0.FalseIdx = 3

	monkey1 := NewMonkey(54, 65, 75, 74)
	monkey1.Op = *NewOperation("old + 6")
	monkey1.Divisor = 19
	monkey1.TrueIdx = 2
	monkey1.FalseIdx = 0

	monkey2 := NewMonkey(79, 60, 97)
	monkey2.Op = *NewOperation("old * old")
	monkey2.Divisor = 13
	monkey2.TrueIdx = 1
	monkey2.FalseIdx = 3

	monkey3 := NewMonkey(74)
	monkey3.Op = *NewOperation("old + 3")
	monkey3.Divisor = 17
	monkey3.TrueIdx = 0
	monkey3.FalseIdx = 1

	friends := []*Monkey{
		monkey0,
		monkey1,
		monkey2,
		monkey3,
	}

	got := monkey0.inspectItemsPartOne(friends)
	expected := friends

	if len(got) != len(expected) {
		t.Errorf("inspectItems() = %v, want %v", got, expected)
	}

	for i := 0; i < 20; i++ {
		monkey0.inspectItemsPartOne(friends)
		monkey1.inspectItemsPartOne(friends)
		monkey2.inspectItemsPartOne(friends)
		monkey3.inspectItemsPartOne(friends)
	}

	gotRes := getTwoMostInspected(friends)
	expectedRes := 10605

	if gotRes != expectedRes {
		t.Errorf("inspectItems() = %v, want %v", gotRes, expectedRes)
	}
}

func TestParseTrueIdx(t *testing.T) {
	input := "If true: throw to monkey 2"
	got := parseTrueIdx(input)
	expected := 2
	if got != expected {
		t.Errorf("parseTrueIdx() = %v, want %v", got, expected)
	}

}

func TestParseFalseIdx(t *testing.T) {
	input := "If false: throw to monkey 3"
	got := parseFalseIdx(input)
	expected := 3
	if got != expected {
		t.Errorf("parseFalseIdx() = %v, want %v", got, expected)
	}

}

func TestParseAllMonkeys(t *testing.T) {
	input := utils.SplitInput(sample)
	var monkeys []*Monkey
	parseAllMonkeys(input, monkeys)
}
