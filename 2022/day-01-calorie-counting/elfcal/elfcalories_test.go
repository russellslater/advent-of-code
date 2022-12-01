package elfcal_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-01-calorie-counting/elfcal"
)

func TestExample(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		input []elfcal.ElfCalories
		top   int
		want  int
	}{
		{
			"Advent of Code Part One Example",
			[]elfcal.ElfCalories{
				{1000, 2000, 3000},
				{4000},
				{5000, 6000},
				{7000, 8000, 9000},
				{10000},
			},
			1,
			24000,
		},
		{
			"Advent of Code Part Two Example",
			[]elfcal.ElfCalories{
				{1000, 2000, 3000},
				{4000},
				{5000, 6000},
				{7000, 8000, 9000},
				{10000},
			},
			3,
			45000,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := elfcal.TopElfTotalCalories(tc.input, tc.top)

			is.Equal(got, tc.want)
		})
	}
}
