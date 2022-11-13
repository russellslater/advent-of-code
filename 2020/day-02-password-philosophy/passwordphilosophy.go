package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	entries := getEntries("2020/day-02-password-philosophy/input.txt")
	validCount := 0
	for _, p := range entries {
		if p.isValid() {
			validCount++
		}
	}
	fmt.Printf("Valid passwords: %d\n", validCount)
}

type policyEntry struct {
	min      int
	max      int
	char     rune
	password string
}

func (p *policyEntry) isValid() bool {
	count := 0
	for _, c := range p.password {
		if c == p.char {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func getEntries(filename string) []*policyEntry {
	input := []*policyEntry{}
	for _, line := range util.LoadInput(filename) {
		r := regexp.MustCompile(`(?P<min>\d{1,2})-(?P<max>\d{1,2}) (?P<char>[a-z]{1}): (?P<password>\w+)`)
		match := r.FindStringSubmatch(line)

		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		char := []rune(match[3])[0]

		input = append(input, &policyEntry{
			min:      min,
			max:      max,
			char:     char,
			password: match[4],
		})
	}
	return input
}
