package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	entries := getEntries("2020/day-02-password-philosophy/input.txt")
	firstRuleValidCount, secondRuleValidCount := 0, 0
	for _, p := range entries {
		if p.isValid(simpleMinMaxPolicyRule) {
			firstRuleValidCount++
		}
		if p.isValid(runePositionRule) {
			secondRuleValidCount++
		}
	}
	fmt.Printf("%d valid passwords according to first rule\n", firstRuleValidCount)
	fmt.Printf("%d valid passwords according to second rule\n", secondRuleValidCount)
}

type passwordPolicyRule func(p *policyEntry) bool

type policyEntry struct {
	min      int
	max      int
	char     rune
	password string
}

func (p *policyEntry) isValid(policy passwordPolicyRule) bool {
	return policy(p)
}

func simpleMinMaxPolicyRule(p *policyEntry) bool {
	count := 0
	for _, c := range p.password {
		if c == p.char {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func runePositionRule(p *policyEntry) bool {
	if p.min <= 0 || p.max > len(p.password) {
		return false
	}
	minMatch := p.password[p.min-1] == byte(p.char)
	maxMatch := p.password[p.max-1] == byte(p.char)
	return (minMatch && !maxMatch) || (!minMatch && maxMatch)
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
