package crane_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-05-supply-stacks/crane"
)

func TestBuildMoves(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		input []string
		want  []crane.Move
	}{
		{
			"Standard moves",
			[]string{
				"move 11 from 2 to 1",
				"move 7 from 7 to 3",
				"move 0 from 0 to 0",
				"move 4 from 39 to 100",
			},
			[]crane.Move{{11, 2, 1}, {7, 7, 3}, {0, 0, 0}, {4, 39, 100}},
		},
		{
			"Garbage input",
			[]string{
				"garbage 1! -input\\/ x09",
			},
			[]crane.Move{{0, 0, 0}},
		},
		{
			"Empty input",
			[]string{},
			[]crane.Move{},
		},
		{
			"Nil input",
			nil,
			[]crane.Move{},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := crane.BuildMoves(tc.input)
			is.Equal(got, tc.want)
		})
	}
}
