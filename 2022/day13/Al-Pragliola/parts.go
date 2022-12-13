package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type List struct {
	Left  *Node
	Right *Node
	Index int
}

type Node struct {
	Value  *int
	Nodes  *[]Node
	Parent *Node
}

func Compare(a *Node, b *Node) (bool, bool) {
	fmt.Printf("Comparing %s to %s\n", a.String(), b.String())

	if b.Nodes == nil && b.Value == nil {
		fmt.Printf("Right side ran out of items, so inputs are not in the right order\n")
		return false, true
	}

	if a.Nodes == nil && a.Value == nil {
		fmt.Printf("Left side ran out of items, so inputs are in the right order\n")
		return true, true
	}

	if a.Nodes != nil && b.Nodes == nil {
		b.Nodes = &[]Node{{
			Value: b.Value,
		}}
	}

	if a.Nodes == nil && b.Nodes != nil {
		a.Nodes = &[]Node{{
			Value: a.Value,
		}}
	}

	if a.Value != nil && b.Value != nil {
		if *a.Value > *b.Value {
			fmt.Printf("Right side is smaller, so inputs are not in the right order\n")
			return false, true
		}

		if *a.Value < *b.Value {
			fmt.Printf("Left side is smaller, so inputs are in the right order\n")
			return true, true
		}
	}

	if a.Nodes != nil && b.Nodes != nil {
		for i := 0; i < len(*a.Nodes); i++ {
			nextNode := Node{}

			if i < len(*b.Nodes) {
				nextNode = (*b.Nodes)[i]
			}

			res, brk := Compare(&(*a.Nodes)[i], &nextNode)
			if brk {
				return res, true
			}
		}

		if len(*a.Nodes) < len(*b.Nodes) {
			fmt.Printf("Left side ran out of items, so inputs are in the right order\n")
			return true, true
		}
	}

	return true, false
}

func (n *Node) String() string {
	if n.Value != nil {
		return strconv.Itoa(*n.Value)
	}

	var result string

	if n.Nodes == nil {
		return ""
	}

	for i, node := range *n.Nodes {
		if node.Value != nil {
			if i > 0 {
				result += ","
			}
			result += strconv.Itoa(*node.Value)
		} else {
			if i > 0 {
				result += ","
			}
			result += fmt.Sprintf("[%s]", node.String())
		}
	}

	return result
}

func ParseInput(in string, i int, parent *Node) *Node {
	re := regexp.MustCompile("[0-9]+")

	if i == len(in) {
		return parent
	}

	var node Node

	if in[i] == '[' {
		node.Parent = parent
		ParseInput(in, i+1, &node)
	} else if in[i] >= '0' && in[i] <= '9' {
		num := re.FindAllStringIndex(in[i:], 1)[0]
		val, _ := strconv.Atoi(in[i+num[0] : i+num[1]])
		node.Value = &val
		node.Parent = parent
		if parent.Nodes == nil {
			parent.Nodes = &[]Node{}
		}
		*parent.Nodes = append(*parent.Nodes, node)
		ParseInput(in, i+num[1], parent)
	} else if in[i] == ',' {
		ParseInput(in, i+1, parent)
	} else {
		if parent.Parent == nil {
			return parent
		}

		if parent.Parent.Nodes == nil {
			parent.Parent.Nodes = &[]Node{}
		}

		*parent.Parent.Nodes = append(*parent.Parent.Nodes, *parent)

		ParseInput(in, i+1, parent.Parent)
	}

	return parent
}

func PartOne(input []string) string {
	var result int
	var lists []List

	idx := 1

	for i := 0; i < len(input); i += 3 {
		list := List{}
		list.Left = &Node{}
		list.Right = &Node{}

		_ = ParseInput(input[i], 1, list.Left)
		_ = ParseInput(input[i+1], 1, list.Right)
		list.Index = idx

		lists = append(lists, list)

		idx++
	}

	for i := 0; i < len(lists); i++ {
		fmt.Printf("idx %d\n", lists[i].Index)
		res, _ := Compare(lists[i].Left, lists[i].Right)
		if res {
			result += lists[i].Index
		}
	}

	return strconv.Itoa(result)
}

func PartTwo(input []string) string {
	var list []Node
	var firstIdx int
	var secondIdx int

	for i := 0; i < len(input); i += 3 {
		nodeOne := Node{}
		nodeTwo := Node{}
		_ = ParseInput(input[i], 1, &nodeOne)
		_ = ParseInput(input[i+1], 1, &nodeTwo)

		list = append(list, nodeOne)
		list = append(list, nodeTwo)
	}

	firstDivider := Node{}
	secondDivider := Node{}

	_ = ParseInput("[[2]]", 1, &firstDivider)
	_ = ParseInput("[[6]]", 1, &secondDivider)

	list = append(list, firstDivider)
	list = append(list, secondDivider)

	sort.Slice(list, func(i, j int) bool {
		res, _ := Compare(&list[i], &list[j])
		return res
	})

	for i := 0; i < len(list); i++ {
		if fmt.Sprintf("[%s]", list[i].String()) == "[[2]]" {
			firstIdx = i + 1
		}

		if fmt.Sprintf("[%s]", list[i].String()) == "[[6]]" {
			secondIdx = i + 1
		}
	}

	return strconv.Itoa(firstIdx * secondIdx)
}
