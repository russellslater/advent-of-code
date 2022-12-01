package main

import (
	"fmt"
	"testing"

	"github.com/russellslater/advent-of-code/2022/day-01-calorie-counting/elfcal"
)

func BenchmarkSolution(b *testing.B) {
	input := getTransformedInput("input.txt")

	tt := []struct {
		top int
	}{
		{0},
		{1},
		{50},
		{100},
	}

	for _, tc := range tt {
		b.Run(fmt.Sprintf("top_%d", tc.top), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				elfcal.TopElfTotalCalories(input, tc.top)
			}
		})
	}
}
