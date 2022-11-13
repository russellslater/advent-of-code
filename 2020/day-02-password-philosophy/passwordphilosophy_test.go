package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestPolicyEntryIsValid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name        string
		policyEntry policyEntry
		want        bool
	}{
		{
			"Too few y characters",
			policyEntry{min: 2, max: 5, char: 'y', password: "dhaixy"},
			false,
		},
		{
			"Too many a characters",
			policyEntry{min: 0, max: 1, char: 'a', password: "aaa"},
			false,
		},
		{
			"Just right",
			policyEntry{min: 5, max: 5, char: 'b', password: "rbibbjbdbkz"},
			true,
		},
		{
			"Empty password",
			policyEntry{min: 0, max: 0, char: 'x', password: ""},
			true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.policyEntry.isValid()

			is.Equal(got, tc.want)
		})
	}
}
