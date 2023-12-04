package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2023/day-04-scratchcards/scratch"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2023/day-04-scratchcards/input.txt"

	lines := util.LoadInput(filename)

	cards := []*scratch.Card{}

	for i, line := range lines {
		card := scratch.NewCard(line)
		card.Ordinal = i + 1
		cards = append(cards, card)
	}

	totalScore := 0
	for _, card := range cards {
		totalScore += card.Score()
	}

	fmt.Printf("Part One Answer: %d\n", totalScore)

	for i, card := range cards {
		matches := card.Matches()
		// assign the number of copies to the next cards
		for j := 1; j <= matches && i+j < len(cards); j++ {
			cards[i+j].Copies += card.Copies
		}
	}

	totalCards := 0
	for _, card := range cards {
		totalCards += card.Copies
	}

	fmt.Printf("Part Two Answer: %d\n", totalCards)
}
