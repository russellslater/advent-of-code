package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-02-rock-paper-scissors/rkpapsiz"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-02-rock-paper-scissors/input.txt"
	rounds := getTransformedInput(filename)

	score := scoreRounds(rounds, rkpapsiz.PlayInResponseScorer)
	fmt.Printf("Part One Answer: %d\n", score)

	score = scoreRounds(rounds, rkpapsiz.EndRoundAsRequiredScorer)
	fmt.Printf("Part Two Answer: %d\n", score)
}

func getTransformedInput(filename string) []*rkpapsiz.Round {
	input := []*rkpapsiz.Round{}
	for _, line := range util.LoadInput(filename) {
		input = append(input, &rkpapsiz.Round{Opponent: rune(line[0]), You: rune(line[2])})
	}
	return input
}

func scoreRounds(rounds []*rkpapsiz.Round, scorer rkpapsiz.StrategyScorer) int {
	totalScore := 0
	for _, r := range rounds {
		totalScore += r.Score(scorer)
	}
	return totalScore
}
