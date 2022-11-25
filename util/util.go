package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	CheckErr(err)

	return SplitInput(string(content))
}

func SplitInput(content string) []string {
	stringContent := strings.TrimSpace(string(content))
	return strings.Split(stringContent, "\n")
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Converter[T, V any] func(T) (V, error)

func Convert[T, V any](value T, converter Converter[T, V]) (V, error) {
	return converter(value)
}

func StringSliceToIntSliceConverter(arr []string) ([]int, error) {
	converted := []int{}

	for _, v := range arr {
		res, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		converted = append(converted, res)
	}

	return converted, nil
}
