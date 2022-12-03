package rksk_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-03-rucksack-reorganization/rksk"
)

func TestReoccuringItem(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name     string
		rucksack rksk.Rucksack
		want     rksk.Item
	}{
		{
			"AoC Example 1",
			rksk.Rucksack{"vJrwpWtwJgWrhcsFMMfFFhFp"},
			rksk.Item('p'),
		},
		{
			"AoC Example 2",
			rksk.Rucksack{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
			rksk.Item('L'),
		},
		{
			"AoC Example 3",
			rksk.Rucksack{"PmmdzqPrVvPwwTWBwg"},
			rksk.Item('P'),
		},
		{
			"AoC Example 4",
			rksk.Rucksack{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
			rksk.Item('v'),
		},
		{
			"AoC Example 5",
			rksk.Rucksack{"ttgJtRGJQctTZtZT"},
			rksk.Item('t'),
		},
		{
			"AoC Example 6",
			rksk.Rucksack{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			rksk.Item('s'),
		},
		{
			"Empty rucksack",
			rksk.Rucksack{},
			0,
		},
		{
			"No repeating items in rucksack",
			rksk.Rucksack{"abc"},
			0,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.rucksack.ReoccuringItem()
			is.Equal(got, tc.want)
		})
	}
}
