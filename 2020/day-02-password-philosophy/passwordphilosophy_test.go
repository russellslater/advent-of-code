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
		policy      passwordPolicyRule
		want        bool
	}{
		{
			"Too few y characters with mix / max policy rule",
			policyEntry{min: 2, max: 5, char: 'y', password: "dhaixy"},
			simpleMinMaxPolicyRule,
			false,
		},
		{
			"Too many a characters with mix / max policy rule",
			policyEntry{min: 0, max: 1, char: 'a', password: "aaa"},
			simpleMinMaxPolicyRule,
			false,
		},
		{
			"Just right with mix / max policy rule",
			policyEntry{min: 5, max: 5, char: 'b', password: "rbibbjbdbkz"},
			simpleMinMaxPolicyRule,
			true,
		},
		{
			"Empty password with mix / max policy rule",
			policyEntry{min: 0, max: 0, char: 'x', password: ""},
			simpleMinMaxPolicyRule,
			true,
		},
		{
			"1-3 a: abcde is valid with rune position rule",
			policyEntry{min: 1, max: 3, char: 'a', password: "abcde"},
			runePositionRule,
			true,
		},
		{
			"1-3 b: cdefg is invalid with rune position rule",
			policyEntry{min: 1, max: 3, char: 'b', password: "cdefg"},
			runePositionRule,
			false,
		},
		{
			"2-9 c: ccccccccc is invalid with rune position rule",
			policyEntry{min: 2, max: 9, char: 'c', password: "ccccccccc"},
			runePositionRule,
			false,
		},
		{
			"Empty password with rune position rule",
			policyEntry{min: 0, max: 0, char: 'x', password: ""},
			runePositionRule,
			false,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			got := tc.policyEntry.isValid(tc.policy)

			is.Equal(got, tc.want)
		})
	}
}
