package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golangroma/aoc/utils"
)

var sample string = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPartOne(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example",
			input:    sample,
			expected: "95437",
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
			expected: "24933642",
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

func TestElfFileIsParsable(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "OK",
			input:    "14848514 b.txt",
			expected: "true",
		},
		{
			name:     "KO_1",
			input:    "14848514",
			expected: "false",
		},
		{
			name:     "KO_2",
			input:    "b.txt 14848514",
			expected: "false",
		},
		{
			name:     "KO_3",
			input:    "$ cd /",
			expected: "false",
		},
		{
			name:     "KO_4",
			input:    "$ ls",
			expected: "false",
		},
		{
			name:     "KO_5",
			input:    "dir a",
			expected: "false",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var ef ElfFile
			if got := strconv.FormatBool(ef.isParsable(tc.input)); got != tc.expected {
				t.Errorf("TestElfFileIsParsable() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestElfFileParse(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "OK",
			input:    "14848514 b.txt",
			expected: "ElfFile Name: b.txt - Size: 14848514",
		},
		{
			name:     "KO_1",
			input:    "14848514",
			expected: "ElfFile Name:  - Size: 0",
		},
		{
			name:     "KO_2",
			input:    "b.txt 14848514",
			expected: "ElfFile Name:  - Size: 0",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ef := parseElfFile(tc.input)
			got := ef.toString()
			if got != tc.expected {
				t.Errorf("TestElfFileParse() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestElfDirectoryIsParsable(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "OK",
			input:    "dir b",
			expected: "true",
		},
		{
			name:     "KO_1",
			input:    "14848514",
			expected: "false",
		},
		{
			name:     "KO_2",
			input:    "b.txt 14848514",
			expected: "false",
		},
		{
			name:     "KO_3",
			input:    "$ cd /",
			expected: "false",
		},
		{
			name:     "KO_4",
			input:    "$ ls",
			expected: "false",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var ed ElfDirectory
			if got := strconv.FormatBool(ed.isParsable(tc.input)); got != tc.expected {
				t.Errorf("TestElfDirectoryIsParsable() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestElfDirectoryParse(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "OK",
			input:    "dir b",
			expected: "ElfDirectory Name: b",
		},
		{
			name:     "KO_1",
			input:    "14848514",
			expected: "ElfDirectory Name: ",
		},
		{
			name:     "KO_2",
			input:    "b.txt 14848514",
			expected: "ElfDirectory Name: ",
		},
	}

	root := newElfDirectory("/", nil)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ed := parseElfDirectory(tc.input, &root)
			got := ed.toString()
			if got != tc.expected {
				t.Errorf("TestElfDirectoryParse() = %v, want %v", got, tc.expected)
			}
		})
	}

	rootGot := newElfDirectory("/", nil)
	rootExpected := newElfDirectory("/", nil)
	testInput := "dir a"
	aDirGot := parseElfDirectory(testInput, &rootGot)
	aDirExpected := newElfDirectory("a", &rootExpected)

	if aDirGot.name != aDirExpected.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", aDirGot.name, aDirExpected.name)
	}

	if aDirGot.name != aDirExpected.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", aDirGot.name, aDirExpected.name)
	}

	if aDirGot.parent.name != rootGot.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", aDirGot.parent.name, rootGot.name)
	}

	if len(aDirGot.children) != 0 {
		t.Errorf("TestElfDirectoryParse() = %d, want %d", len(aDirGot.children), 0)
	}

	testInput = "14848514 b.txt"
	fileGot := parseElfFile(testInput)
	fileExpected := newElfFile("b.txt", 14848514)

	if fileGot != fileExpected {
		t.Errorf("TestElfFileParse() = %s, want %s", fileGot.toString(), fileExpected.toString())
	}

	aDirGot.addFile(fileGot)

	testInput = "dir c"
	cDirGot := parseElfDirectory(testInput, aDirGot)
	cDirExpected := newElfDirectory("c", aDirGot)

	if cDirGot.name != cDirExpected.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", cDirGot.name, cDirExpected.name)
	}

	if cDirGot.name != cDirExpected.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", cDirGot.name, cDirExpected.name)
	}

	if cDirGot.parent.name != aDirGot.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", cDirGot.parent.name, aDirGot.name)
	}

	if len(aDirGot.children) != 1 {
		t.Errorf("TestElfDirectoryParse() = %d, want %d", len(aDirGot.children), 1)
	}

	if aDirGot.children["c"].name != cDirGot.name {
		t.Errorf("TestElfDirectoryParse() = %s, want %s", aDirGot.children["c"].name, cDirGot.name)
	}
}

func TestBuildFileSystem(t *testing.T) {
	input := utils.SplitInput(sample)
	root := buildFileSystem(input)
	totalSize := root.size()
	lts := getLimitedTotalDirectorySize(&root)
	fmt.Println(totalSize)
	fmt.Println(lts)
	diskOccupation := root.size()
	toFree := diskOccupation - (diskSize - neededFreeSpace)
	result := findSmallestDirectoryBySize(&root, diskOccupation, toFree)
	fmt.Println(result)
}
