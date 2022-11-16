package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2020/day-04-passport-processing/input.txt"
	passports := getInputPassports(filename)
	count := validPassportCount(passports)
	fmt.Printf("Valid password count: %v\n", count)
}

var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // excludes 'cid'

type passport struct {
	lines []string
}

func (p *passport) hasRequiredFields() bool {
	content := strings.Join(p.lines, " ")
	matchedAll := true
	for _, f := range reqFields {
		if !strings.Contains(content, f+":") {
			matchedAll = false
			break
		}
	}
	return matchedAll
}

func getInputPassports(filename string) []*passport {
	input := []*passport{}

	pl := []string{}
	for _, line := range util.LoadInput(filename) {
		if line == "" {
			input = append(input, &passport{pl})
			pl = []string{}
		}
		pl = append(pl, line)
	}
	input = append(input, &passport{pl})

	return input
}

func validPassportCount(passports []*passport) int {
	count := 0
	for _, p := range passports {
		if p.hasRequiredFields() {
			count++
		}
	}
	return count
}
