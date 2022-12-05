package crane

import (
	"testing"

	"github.com/matryer/is"
)

func TestMoveCrates(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name           string
		handleMultiple bool
		stacks         [][]string
		moves          []Move

		want []string
	}{
		{
			"AoC example stacks",
			false,
			[][]string{
				{"N", "Z"},
				{"D", "C", "M"},
				{"P"},
			},
			[]Move{
				{1, 2, 1},
				{3, 1, 3},
				{2, 2, 1},
				{1, 1, 2},
			},
			[]string{
				"DCP",
				" CZ",
				"M Z",
				"CMZ",
			},
		},
		{
			"AoC provided stacks",
			false,
			[][]string{
				{"T", "F", "V", "Z", "C", "W", "S", "Q"},
				{"B", "R", "Q"},
				{"S", "M", "P", "Q", "T", "Z", "B"},
				{"H", "Q", "R", "F", "V", "D"},
				{"P", "T", "S", "B", "D", "L", "G", "J"},
				{"Z", "T", "R", "W"},
				{"J", "R", "F", "S", "N", "M", "Q", "H"},
				{"W", "H", "F", "N", "R"},
				{"B", "P", "R", "Q", "T", "Z", "J"},
			},
			[]Move{
				{1, 1, 2},
				{5, 8, 9},
				{3, 3, 8},
				{5, 4, 7},
				{2, 6, 6}, // place back onto same stack
			},
			[]string{
				"FTSHPZJWB",
				"FTSHPZJ R",
				"FTQHPZJPR",
				"FTQDPZVPR",
				"FTQDPTVPR",
			},
		},
		{
			"AoC example stacks with multiple crates moved at once",
			true,
			[][]string{
				{"N", "Z"},
				{"D", "C", "M"},
				{"P"},
			},
			[]Move{
				{1, 2, 1},
				{3, 1, 3},
				{2, 2, 1},
				{1, 1, 2},
			},
			[]string{
				"DCP",
				" CD",
				"C D",
				"MCD",
			},
		},
		{
			"AoC provided stacks with multiple crates moved at once",
			true,
			[][]string{
				{"T", "F", "V", "Z", "C", "W", "S", "Q"},
				{"B", "R", "Q"},
				{"S", "M", "P", "Q", "T", "Z", "B"},
				{"H", "Q", "R", "F", "V", "D"},
				{"P", "T", "S", "B", "D", "L", "G", "J"},
				{"Z", "T", "R", "W"},
				{"J", "R", "F", "S", "N", "M", "Q", "H"},
				{"W", "H", "F", "N", "R"},
				{"B", "P", "R", "Q", "T", "Z", "J"},
			},
			[]Move{
				{1, 1, 2},
				{5, 8, 9},
				{3, 3, 8},
				{5, 4, 7},
				{2, 6, 6}, // place back onto same stack
			},
			[]string{
				"FTSHPZJWB",
				"FTSHPZJ W",
				"FTQHPZJSW",
				"FTQDPZHSW",
				"FTQDPZHSW", // no change
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			crane := &Crane{handleMultiple: tc.handleMultiple}

			stacksCopy := make([][]string, len(tc.stacks))
			copy(stacksCopy, tc.stacks)

			// assert each move results in expected output
			for i, mv := range tc.moves {
				stacks := crane.moveCrates(tc.stacks, mv)
				got := crane.result(stacks)
				is.Equal(got, tc.want[i])
			}

			// ensure Operate lands on same end result
			got := crane.Operate(stacksCopy, tc.moves)
			is.Equal(got, tc.want[len(tc.want)-1])
		})
	}
}
