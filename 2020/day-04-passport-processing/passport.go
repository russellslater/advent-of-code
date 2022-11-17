package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2020/day-04-passport-processing/input.txt"
	passports := getInputPassports(filename)

	count := validPassportCount(passports, requiredFieldsOnlyPolicy)
	fmt.Printf("Passports with required fields only: %v\n", count)

	count = validPassportCount(passports, validRequiredFieldsPolicy)
	fmt.Printf("Passports with valid required fields: %v\n", count)
}

var reqFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // excludes 'cid'
var fieldRules = []string{
	"\\b(byr:19[2-9][0-9]|200[0-2])\\b",
	"\\b(iyr:20((1[0-9])|20))\\b",
	"\\b(eyr:20((2[0-9])|30))\\b",
	"\\b(hgt:1([5-8][0-9]|9[0-3])cm)|(hgt:(59|6[0-9]|7[0-6])in)\\b",
	"\\b(hcl:#[a-f0-9]{6})\\b",
	"\\b(ecl:(amb|blu|brn|gry|grn|hzl|oth))\\b",
	"\\b(pid:[0-9]{9})\\b",
}

func requiredFieldsOnlyPolicy(passportContent string) bool {
	matchedAll := true
	for _, f := range reqFields {
		if !strings.Contains(passportContent, f+":") {
			matchedAll = false
			break
		}
	}
	return matchedAll
}

func validRequiredFieldsPolicy(passportContent string) bool {
	matchedAll := true
	for _, r := range fieldRules {
		if m, _ := regexp.MatchString(r, passportContent); !m {
			matchedAll = false
			break
		}
	}
	return matchedAll
}

type passport struct {
	lines []string
}

func (p *passport) isValid(policy func(content string) bool) bool {
	content := strings.Join(p.lines, " ")
	return policy(content)
}

func getInputPassports(filename string) []*passport {
	input := []*passport{}

	pl := []string{}
	for _, line := range util.LoadInput(filename) {
		if line == "" {
			input = append(input, &passport{pl})
			pl = []string{}
		} else {
			pl = append(pl, line)
		}
	}
	input = append(input, &passport{pl})

	return input
}

func validPassportCount(passports []*passport, policy func(content string) bool) int {
	count := 0
	for _, p := range passports {
		if p.isValid(policy) {
			count++
		}
	}
	return count
}
