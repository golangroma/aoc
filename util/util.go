package util

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	CheckErr(err)

	return strings.Split(string(content), "\n")
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
