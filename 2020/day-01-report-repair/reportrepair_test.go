package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestFind2020PairSumProduct(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		target int
		input  []int
		want   int
	}{
		{
			"Zero array",
			2020,
			[]int{0},
			0,
		},
		{
			"Empty array",
			2020,
			[]int{},
			0,
		},
		{
			"Simple input",
			2020,
			[]int{1840, 180},
			331_200,
		},
		{
			"Example input",
			2020,
			[]int{1721, 979, 366, 299, 675, 1456},
			514_579,
		},
		{
			"Input where two numbers do not sum up to target number",
			2020,
			[]int{1720, 979, 366, 299, 675, 1456},
			0,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := FindPairSumProduct(tc.target, tc.input)
			is.Equal(got, tc.want) // product not expected
		})
	}
}

func TestFind2020TrioSumProduct(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		target int
		input  []int
		want   int
	}{
		{
			"Zero array",
			2020,
			[]int{0},
			0,
		},
		{
			"Empty array",
			2020,
			[]int{},
			0,
		},
		{
			"Simple input",
			2020,
			[]int{1840, 159, 21},
			6_143_760,
		},
		{
			"Example input",
			2020,
			[]int{1721, 979, 366, 299, 675, 1456},
			241_861_950,
		},
		{
			"Input where two numbers do not sum up to target number",
			2020,
			[]int{1721, 999, 366, 299, 675, 1456},
			0,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := FindTrioSumProduct(tc.target, tc.input)
			is.Equal(got, tc.want) // product not expected
		})
	}
}
