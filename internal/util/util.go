package util

import (
	"bufio"
	"log"
	"os"
)

func LoadInput(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer f.Close()

	input := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func Prepend(arr []string, items []string) []string {
	for _, item := range items {
		arr = append([]string{item}, arr...)
	}
	return arr
}

func Reverse(arr []string) []string {
	reversed := make([]string, len(arr))
	for i, item := range arr {
		reversed[len(arr)-1-i] = item
	}
	return reversed
}

func Unique(arr []rune) bool {
	if len(arr) == 0 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range arr {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
