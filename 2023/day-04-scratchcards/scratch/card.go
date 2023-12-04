package scratch

import (
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Ordinal        int
	WinningNumbers []int
	MyNumbers      []int
	Copies         int
}

func NewCard(line string) *Card {
	parts := strings.Split(line, "|")
	winningNumbers := parseNumbers(parts[0])
	myNumbers := parseNumbers(parts[1])

	return &Card{
		WinningNumbers: winningNumbers,
		MyNumbers:      myNumbers,
		Copies:         1,
	}
}

func (c Card) Score() int {
	score := 0
	multiplier := 1
	for _, num := range c.MyNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			score = multiplier
			multiplier *= 2
		}
	}
	return score
}

func (c Card) Matches() int {
	matches := 0
	for _, num := range c.MyNumbers {
		if slices.Contains(c.WinningNumbers, num) {
			matches++
		}
	}
	return matches
}

func parseNumbers(s string) []int {
	var nums []int
	fields := strings.Fields(s)
	for _, field := range fields {
		// throw away instances of "Card" and "XX:"
		if num, err := strconv.Atoi(field); err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}
