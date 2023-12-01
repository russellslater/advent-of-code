package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/2023/day-01-trebuchet/digit"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2023/day-01-trebuchet/input.txt"
	lines := util.LoadInput(filename)

	sum := 0
	for _, line := range lines {
		first, last := firstAndLastNumericDigit(line)
		sum += util.MustAtoi(fmt.Sprintf("%d%d", first.Value, last.Value))
	}

	fmt.Printf("Part One Answer: %d\n", sum)

	sum = 0
	for _, line := range lines {
		first, last := firstAndLastNumericDigit(line)
		firstSpelled, lastSpelled := findFirstAndLastSpelledDigits(line)

		if !firstSpelled.IsEmpty() && (first.IsEmpty() || !firstSpelled.HasLargerIndex(first)) {
			first = firstSpelled
		}
		if !lastSpelled.IsEmpty() && (last.IsEmpty() || lastSpelled.HasLargerIndex(last)) {
			last = lastSpelled
		}

		sum += util.MustAtoi(fmt.Sprintf("%d%d", first.Value, last.Value))
	}

	fmt.Printf("Part Two Answer: %d\n", sum)
}

func firstAndLastNumericDigit(str string) (digit.Digit, digit.Digit) {
	first, last := digit.EmptyDigit, digit.EmptyDigit

	for i, c := range str {
		if c >= '0' && c <= '9' {
			num := util.MustAtoi(string(c))
			if first.IsEmpty() {
				first = digit.Digit{Idx: i, Value: num}
			}
			last = digit.Digit{Idx: i, Value: num}
		}
	}

	return first, last
}

func findFirstAndLastSpelledDigits(str string) (digit.Digit, digit.Digit) {
	firstSpelledNumber := digit.EmptyDigit
	for i, d := range digit.SpelledDigits {
		idx := strings.Index(str, d)
		if idx != -1 && (firstSpelledNumber.IsEmpty() || idx < firstSpelledNumber.Idx) {
			firstSpelledNumber = digit.Digit{Idx: idx, Value: i + 1}
		}
	}

	lastSpelledNumber := digit.EmptyDigit
	for i, d := range digit.SpelledDigits {
		idx := strings.LastIndex(str, d)
		if idx != -1 && (lastSpelledNumber.IsEmpty() || idx > lastSpelledNumber.Idx) {
			lastSpelledNumber = digit.Digit{Idx: idx, Value: i + 1}
		}
	}

	return firstSpelledNumber, lastSpelledNumber
}
