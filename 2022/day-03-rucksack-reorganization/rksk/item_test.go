package rksk_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/russellslater/advent-of-code/2022/day-03-rucksack-reorganization/rksk"
)

func TestItemScore(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	want := 1 // starting priority

	for i := 'a'; i <= 'z'; i++ {
		got := rksk.Item(i).Score()
		is.Equal(got, want)
		want++
	}

	for i := 'A'; i <= 'Z'; i++ {
		got := rksk.Item(i).Score()
		is.Equal(got, want)
		want++
	}
}
