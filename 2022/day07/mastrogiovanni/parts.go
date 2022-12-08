package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CmdType int

const (
	CmdLs CmdType = iota
	CmdCd
	CmdCdUp
	CmdFile
	CmdDir
)

type Command struct {
	Type CmdType
	Size int
	Name string
}

func (c Command) String() string {
	if c.Type == CmdCd {
		return fmt.Sprintf("$ cd %s", c.Name)
	}
	if c.Type == CmdCdUp {
		return "$ cd .."
	}
	if c.Type == CmdLs {
		return "$ ls"
	}
	if c.Type == CmdDir {
		return fmt.Sprintf("dir %s", c.Name)
	}
	if c.Type == CmdFile {
		return fmt.Sprintf("%d %s", c.Size, c.Name)
	}
	return ""
}

type NodeType int

const (
	File NodeType = iota
	Directory
)

type Node struct {
	Type     NodeType
	Name     string
	Parent   *Node
	Children []*Node
	Size     int
}

func (n *Node) appendFile(name string, size int) {
	n.Children = append(n.Children, &Node{
		Type:     File,
		Name:     name,
		Parent:   n,
		Children: nil,
		Size:     size,
	})
}

func (n *Node) appendDirectory(name string) {
	n.Children = append(n.Children, &Node{
		Type:     Directory,
		Name:     name,
		Parent:   n,
		Children: make([]*Node, 0),
		Size:     0,
	})
}

func (n *Node) cd(name string) *Node {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func ParseCommand(input string) Command {
	comps := strings.Split(input, " ")
	if comps[0] == "$" {
		if comps[1] == "cd" {
			if comps[2] == ".." {
				return Command{
					Type: CmdCdUp,
				}
			} else {
				return Command{
					Type: CmdCd,
					Name: comps[2],
				}
			}
		} else if comps[1] == "ls" {
			return Command{
				Type: CmdLs,
			}

		}
	} else {
		if comps[0] == "dir" {
			return Command{
				Type: CmdDir,
				Name: comps[1],
			}
		} else {
			size, _ := strconv.Atoi(comps[0])
			return Command{
				Type: CmdFile,
				Size: size,
				Name: comps[1],
			}
		}
	}
	return Command{}
}

func ParseCommands(inputs []string) []Command {
	result := make([]Command, 0)
	for _, input := range inputs {
		result = append(result, ParseCommand(input))
	}
	return result
}

func initTree() *Node {
	return &Node{
		Type:   Directory,
		Parent: nil,
		Children: []*Node{
			&Node{
				Type: Directory,
				Name: "/",
			},
		},
	}
}

func ParseTree(commands []Command, root *Node) *Node {
	for i := 0; i < len(commands); i++ {
		command := commands[i]
		switch command.Type {
		case CmdCd:
			root = root.cd(command.Name)
		case CmdCdUp:
			root = root.Parent
		case CmdLs:
			j := i + 1
			for ; j < len(commands); j++ {
				nextCommand := commands[j]
				if nextCommand.Type == CmdDir {
					root.appendDirectory(nextCommand.Name)
				} else if nextCommand.Type == CmdFile {
					root.appendFile(nextCommand.Name, nextCommand.Size)
				} else {
					i = j - 1
					break
				}
			}
			if j == len(commands) {
				for root.Parent != nil {
					root = root.Parent
				}
				return root
			}
		}
	}
	return root
}

func SizeDirectories(root *Node) {
	if root.Type == Directory {
		tot := 0
		for _, child := range root.Children {
			if child.Type == Directory {
				SizeDirectories(child)
			}
			tot += child.Size
		}
		root.Size = tot
	}
}

func Dump(node *Node, tab int) {
	for i := 0; i < tab; i++ {
		fmt.Printf("  ")
	}
	if node.Type == Directory {
		fmt.Printf("- %s (dir, size: %d)\n", node.Name, node.Size)
	} else if node.Type == File {
		fmt.Printf("- %s (file, size: %d)\n", node.Name, node.Size)
	}
	for _, child := range node.Children {
		Dump(child, tab+1)
	}
}

func SumSmallDirectories(root *Node) int {
	tot := 0
	for _, child := range root.Children {
		if child.Type == Directory {
			if child.Size <= 100000 {
				tot += child.Size
			}
			tot += SumSmallDirectories(child)
		}
	}
	return tot
}

func BestDirectoryToRemove(root *Node) int {
	diskOccupation := root.Children[0].Size
	toFree := diskOccupation - (70000000 - 30000000)
	return CompareDirectory(root, diskOccupation, toFree)
}

func CompareDirectory(node *Node, best int, freeSpace int) int {
	for _, child := range node.Children {
		if child.Type == Directory {
			if child.Size >= freeSpace {
				if child.Size < best {
					best = child.Size
				}
			}
			best = CompareDirectory(child, best, freeSpace)
		}
	}
	return best
}

func PartOne(input []string) string {
	commands := ParseCommands(input)
	root := initTree()
	tree := ParseTree(commands, root)
	SizeDirectories(tree)
	tot := SumSmallDirectories(root)
	return fmt.Sprintf("%d", tot)
}

func PartTwo(input []string) string {
	commands := ParseCommands(input)
	root := initTree()
	tree := ParseTree(commands, root)
	SizeDirectories(tree)
	// Dump(tree, 0)
	tot := BestDirectoryToRemove(root)
	return fmt.Sprintf("%d", tot)
}
