package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/russellslater/advent-of-code/2015/day-07-some-assembly-required/sig"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-07-some-assembly-required/input.txt"
	statements := getTransformedInput(filename)

	signal := calcSignal(statements, "a")
	fmt.Printf("Part One Answer: %d\n", signal)

	statements = getTransformedInput(filename)

	for _, statement := range statements {
		if statement.Output == "b" {
			statement.InputOne = signal
			break
		}
	}

	signal = calcSignal(statements, "a")
	fmt.Printf("Part Two Answer: %d\n", signal)
}

func getTransformedInput(filename string) []*sig.Statement {
	passThroughRegexp := regexp.MustCompile(`^(\w+) -> (\w+)$`)
	shiftRegexp := regexp.MustCompile(`^(\w+) (RSHIFT|LSHIFT) (\d+) -> (\w+)$`)
	andOrRegexp := regexp.MustCompile(`^(\w+) (AND|OR) (\w+) -> (\w+)$`)
	notRegexp := regexp.MustCompile(`^NOT (\w+) -> (\w+)$`)

	statements := []*sig.Statement{}
	for _, line := range util.LoadInput(filename) {
		var s *sig.Statement
		if passThroughRegexp.MatchString(line) {
			match := passThroughRegexp.FindStringSubmatch(line)
			s = sig.NewPassThroughStatement(parsePossibleNum(match[1]), match[2])
		}
		if notRegexp.MatchString(line) {
			match := notRegexp.FindStringSubmatch(line)
			s = sig.NewNotStatement(parsePossibleNum(match[1]), match[2])
		}
		if andOrRegexp.MatchString(line) {
			match := andOrRegexp.FindStringSubmatch(line)
			if match[2] == "AND" {
				s = sig.NewAndStatement(parsePossibleNum(match[1]), parsePossibleNum(match[3]), match[4])
			} else {
				s = sig.NewOrStatement(parsePossibleNum(match[1]), parsePossibleNum(match[3]), match[4])
			}
		}
		if shiftRegexp.MatchString(line) {
			match := shiftRegexp.FindStringSubmatch(line)
			if match[2] == "RSHIFT" {
				s = sig.NewRShiftStatement(parsePossibleNum(match[1]), parsePossibleNum(match[3]), match[4])
			} else {
				s = sig.NewLShiftStatement(parsePossibleNum(match[1]), parsePossibleNum(match[3]), match[4])
			}
		}

		if s != nil {
			statements = append(statements, s)
		}
	}
	return statements
}

func parsePossibleNum(input string) any {
	if num, err := strconv.Atoi(input); err == nil {
		return uint16(num)
	}
	return input
}

func calcSignal(statements []*sig.Statement, targetWire string) interface{} {
	store := map[string]uint16{}

	// Rip through the statements until all are operable
	for {
		operableCount := 0
		for _, s := range statements {
			if s.IsOperable() {
				operableCount++
				if res, ok := s.Execute(); ok {
					store[s.Output] = res
				}
			} else {
				s.Inflate(store)
			}
		}

		if operableCount == len(statements) {
			break
		}
	}

	return store[targetWire]
}
