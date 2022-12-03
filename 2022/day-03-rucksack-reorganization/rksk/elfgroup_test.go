package rksk_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-03-rucksack-reorganization/rksk"
)

func TestGroupElves(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name      string
		rucksacks []*rksk.Rucksack
		size      int
		want      []rksk.ElfGroup
	}{
		{
			"Split 6 rucksacks into 2 groups of 3",
			[]*rksk.Rucksack{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			3,
			[]rksk.ElfGroup{
				{
					{"vJrwpWtwJgWrhcsFMMfFFhFp"},
					{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
					{"PmmdzqPrVvPwwTWBwg"},
				},
				{
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
					{"ttgJtRGJQctTZtZT"},
					{"CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
		},
		{
			"Split 6 rucksacks into 6 groups of 1",
			[]*rksk.Rucksack{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			1,
			[]rksk.ElfGroup{
				{
					{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				},
				{
					{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				},
				{
					{"PmmdzqPrVvPwwTWBwg"},
				},
				{
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				},
				{
					{"ttgJtRGJQctTZtZT"},
				},
				{
					{"CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
		},
		{
			"Split 6 rucksacks into 3 groups of 2",
			[]*rksk.Rucksack{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			2,
			[]rksk.ElfGroup{
				{
					{"vJrwpWtwJgWrhcsFMMfFFhFp"},
					{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				},
				{
					{"PmmdzqPrVvPwwTWBwg"},
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				},
				{
					{"ttgJtRGJQctTZtZT"},
					{"CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
		},
		{
			"Size is zero",
			[]*rksk.Rucksack{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			0,
			[]rksk.ElfGroup{},
		},
		{
			"Size is greater than number of rucksacks",
			[]*rksk.Rucksack{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			100,
			[]rksk.ElfGroup{
				{
					{"vJrwpWtwJgWrhcsFMMfFFhFp"},
					{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
					{"PmmdzqPrVvPwwTWBwg"},
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
					{"ttgJtRGJQctTZtZT"},
					{"CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := rksk.GroupElves(tc.rucksacks, tc.size)
			is.Equal(got, tc.want)
		})
	}
}

func TestCommonItem(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name     string
		elfgroup rksk.ElfGroup
		want     rksk.Item
	}{
		{
			"AoC Example: Common item in group of 3 is 'r'",
			rksk.ElfGroup{
				{"vJrwpWtwJgWrhcsFMMfFFhFp"},
				{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"},
				{"PmmdzqPrVvPwwTWBwg"},
			},
			rksk.Item('r'),
		},
		{
			"AoC Example: Common item in group of 3 is 'Z'",
			rksk.ElfGroup{
				{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"},
				{"ttgJtRGJQctTZtZT"},
				{"CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			rksk.Item('Z'),
		},
		{
			"Common item in group of 4 is 'X'",
			rksk.ElfGroup{
				{"aX"},
				{"aX"},
				{"aX"},
				{"bX"},
			},
			rksk.Item('X'),
		},
		{
			"Common item in group of 10 is 'f'",
			rksk.ElfGroup{
				{"abcf"},
				{"def"},
				{"gfhi"},
				{"fjkl"},
				{"mfno"},
				{"pqrf"},
				{"stfu"},
				{"vfwx"},
				{"yzfA"},
				{"BfCD"},
			},
			rksk.Item('f'),
		},
		{
			"No common item in group of 1",
			rksk.ElfGroup{
				{"TTTabc"},
			},
			0,
		},
		{
			"No common item in group of 0",
			rksk.ElfGroup{},
			0,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.elfgroup.CommonItem()
			is.Equal(got, tc.want)
		})
	}
}
