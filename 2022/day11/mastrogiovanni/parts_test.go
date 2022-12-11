package main

import (
	"testing"
)

var (
	m0 = &Monkey{
		worry_levels:    []int{79, 98},
		operation:       func(old int) int { return old * 19 },
		divisor:         23,
		monkey_if_true:  2,
		monkey_if_false: 3,
	}

	m1 = &Monkey{
		worry_levels:    []int{54, 65, 75, 74},
		operation:       func(old int) int { return old + 6 },
		divisor:         19,
		monkey_if_true:  2,
		monkey_if_false: 0,
	}

	m2 = &Monkey{
		worry_levels:    []int{79, 60, 97},
		operation:       func(old int) int { return old * old },
		divisor:         13,
		monkey_if_true:  1,
		monkey_if_false: 3,
	}

	m3 = &Monkey{
		worry_levels:    []int{74},
		operation:       func(old int) int { return old + 3 },
		divisor:         17,
		monkey_if_true:  0,
		monkey_if_false: 1,
	}

	monkeys = []*Monkey{m0, m1, m2, m3}
)

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    []*Monkey
		expected string
	}{
		{
			name:     "example",
			input:    monkeys,
			expected: "10605",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartOne(tc.input); got != tc.expected {
				t.Errorf("PartOne() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	tt := []struct {
		name     string
		input    []*Monkey
		expected string
	}{
		{
			name:     "example",
			input:    monkeys,
			expected: "2713310158",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartTwo(tc.input); got != tc.expected {
				t.Errorf("PartTwo() = %v, want %v", got, tc.expected)
			}
		})
	}
}
