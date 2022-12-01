package main

import (
	"log"
	"os"

	"github.com/golangroma/aoc/cli/internal/challenge"
)

func main() {
	session := os.Getenv("SESSION")
	if err := challenge.Execute(session); err != nil {
		log.Fatal(err)
	}
}
