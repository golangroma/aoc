package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ElfFile struct {
	name string
	size int
}

func newElfFile(name string, size int) ElfFile {
	return ElfFile{name: name, size: size}
}

func (f *ElfFile) isParsable(input string) bool {
	slices := strings.Split(input, " ")

	if len(slices) != 2 {
		return false
	}

	if slices[0][0] >= '1' && slices[0][0] <= '9' {
		return true
	}

	return false
}

func parseElfFile(input string) ElfFile {
	e := ElfFile{}

	if !e.isParsable(input) {
		return e
	}

	slices := strings.Split(input, " ")

	size, _ := strconv.Atoi(slices[0])
	name := slices[1]

	e = newElfFile(name, size)

	return e
}

func (f *ElfFile) toString() string {
	return fmt.Sprintf("ElfFile Name: %s - Size: %d", f.name, f.size)
}
