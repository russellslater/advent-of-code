package asspair_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-04-camp-cleanup/asspair"
)

func TestAssignmentPair(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name            string
		pair            asspair.AssignmentPair
		wantIsContained bool
		wantIsOverlap   bool
	}{
		{
			name:            "Empty AssignmentPair",
			pair:            asspair.AssignmentPair{},
			wantIsContained: true,
			wantIsOverlap:   true,
		},
		{
			name: "Identical ranges in AssignmentPair",
			pair: asspair.AssignmentPair{
				asspair.Range{10, 20},
				asspair.Range{10, 20},
			},
			wantIsContained: true,
			wantIsOverlap:   true,
		},
		{
			name: "Empty range in AssignmentPair",
			pair: asspair.AssignmentPair{
				asspair.Range{1, 10},
				asspair.Range{},
			},
			wantIsContained: false,
			wantIsOverlap:   false,
		},
		{
			name: "AoC Example 1",
			pair: asspair.AssignmentPair{
				asspair.Range{2, 4},
				asspair.Range{6, 8},
			},
			wantIsContained: false,
			wantIsOverlap:   false,
		},
		{
			name: "AoC Example 2",
			pair: asspair.AssignmentPair{
				asspair.Range{2, 3},
				asspair.Range{4, 5},
			},
			wantIsContained: false,
			wantIsOverlap:   false,
		},
		{
			name: "AoC Example 3",
			pair: asspair.AssignmentPair{
				asspair.Range{5, 7},
				asspair.Range{7, 9},
			},
			wantIsContained: false,
			wantIsOverlap:   true,
		},
		{
			name: "AoC Example 4",
			pair: asspair.AssignmentPair{
				asspair.Range{2, 8},
				asspair.Range{3, 7},
			},
			wantIsContained: true,
			wantIsOverlap:   true,
		},
		{
			name: "AoC Example 5",
			pair: asspair.AssignmentPair{
				asspair.Range{6, 6},
				asspair.Range{4, 6},
			},
			wantIsContained: true,
			wantIsOverlap:   true,
		},
		{
			name: "AoC Example 6",
			pair: asspair.AssignmentPair{
				asspair.Range{2, 6},
				asspair.Range{4, 8},
			},
			wantIsContained: false,
			wantIsOverlap:   true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.pair.IsContained()
			is.Equal(got, tc.wantIsContained)

			got = tc.pair.IsOverlap()
			is.Equal(got, tc.wantIsOverlap)
		})
	}
}
