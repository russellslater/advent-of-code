package rkpapsiz_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-02-rock-paper-scissors/rkpapsiz"
)

func TestStrategyScorer(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		scorer rkpapsiz.StrategyScorer
		round  *rkpapsiz.Round
		want   int
	}{
		{
			"Paper beats rock",
			rkpapsiz.PlayInResponseScorer,
			&rkpapsiz.Round{Opponent: 'A', You: 'Y'},
			8,
		},
		{
			"Rock loses to paper",
			rkpapsiz.PlayInResponseScorer,
			&rkpapsiz.Round{Opponent: 'B', You: 'X'},
			1,
		},
		{
			"Scissors draws with scissors",
			rkpapsiz.PlayInResponseScorer,
			&rkpapsiz.Round{Opponent: 'C', You: 'Z'},
			6,
		},
		{
			"Draw against rock with rock",
			rkpapsiz.EndRoundAsRequiredScorer,
			&rkpapsiz.Round{Opponent: 'A', You: 'Y'},
			4,
		},
		{
			"Lose against paper with rock",
			rkpapsiz.EndRoundAsRequiredScorer,
			&rkpapsiz.Round{Opponent: 'B', You: 'X'},
			1,
		},
		{
			"Win against scissors with rock",
			rkpapsiz.EndRoundAsRequiredScorer,
			&rkpapsiz.Round{Opponent: 'C', You: 'Z'},
			7,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.scorer(tc.round)

			is.Equal(got, tc.want)

			got = tc.round.Score(tc.scorer)

			is.Equal(got, tc.want)
		})
	}
}
