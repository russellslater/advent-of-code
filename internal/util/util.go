package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}
