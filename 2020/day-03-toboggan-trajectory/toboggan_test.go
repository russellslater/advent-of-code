package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestPolicyEntryIsValid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name      string
		move      move
		landscape []string
		want      int
	}{
		{
			"Advent of Code 3,1 example",
			move{3, 1},
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			7,
		},
		{
			"Advent of Code 1,1 example",
			move{1, 1},
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			2,
		},
		{
			"Advent of Code 1,2 example",
			move{1, 2},
			[]string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			2,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			l := [][]rune{}
			for _, s := range tc.landscape {
				l = append(l, []rune(s))
			}

			got := treeEncounters(l, tc.move)

			is.Equal(got, tc.want)
		})
	}
}
