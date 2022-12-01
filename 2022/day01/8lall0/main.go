package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Part 1")
	elf, err := part1(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	elves, err := part2(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(elf)
	fmt.Println("Part 2")
	sum := 0
	for _, e := range elves {
		sum += e.Calories
		fmt.Println(e)
	}
	fmt.Printf("Total calories: %d\n", sum)
}
