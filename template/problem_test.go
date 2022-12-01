package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestExample(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name       string
		inputLines []*inputLine
		want       interface{}
	}{
		{
			"Example test",
			[]*inputLine{},
			true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := solveProblem(tc.inputLines)

			is.Equal(got, tc.want)
		})
	}
}
