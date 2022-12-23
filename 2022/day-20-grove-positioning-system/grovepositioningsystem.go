package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-20-grove-positioning-system/input.txt"

	idxNumPairs := getTransformedInput(filename, 1)
	sum := sumGroveCoordinates(idxNumPairs, 1)
	fmt.Printf("Part One Answer: %v\n", sum)

	idxNumPairs = getTransformedInput(filename, 811589153)
	sum = sumGroveCoordinates(idxNumPairs, 10)
	fmt.Printf("Part Two Answer: %v\n", sum)
}

type IdxNumPair struct {
	Idx int
	Num int
}

func getTransformedInput(filename string, decryptionKey int) []IdxNumPair {
	pairs := []IdxNumPair{}
	for i, line := range util.LoadInput(filename) {
		pairs = append(pairs, IdxNumPair{i, util.MustAtoi(line) * decryptionKey})
	}
	return pairs
}

func sumGroveCoordinates(origPairs []IdxNumPair, mixCount int) int {
	mixedPairs := []IdxNumPair{}
	mixedPairs = append(mixedPairs, origPairs...)

	for t := 0; t < mixCount; t++ {
		// Cycle through original pairs in order
		for _, op := range origPairs {
			// Find current index of pair in mixed pairs
			var idx int
			for i, mp := range mixedPairs {
				if op.Idx == mp.Idx && op.Num == mp.Num {
					idx = i
					break
				}
			}

			// Pop off at index
			mixedPairs = append(mixedPairs[:idx], mixedPairs[idx+1:]...)

			// Calculate new position
			pos := (idx + op.Num) % len(mixedPairs)
			if pos <= 0 {
				pos += len(mixedPairs)
			}

			// Pop back in at new position
			mixedPairs = append(mixedPairs[:pos], append([]IdxNumPair{op}, mixedPairs[pos:]...)...)
		}
	}

	var zeroIdx int
	for i, v := range mixedPairs {
		if v.Num == 0 {
			zeroIdx = i
			break
		}
	}

	sum := 0
	for _, p := range []int{1000, 2000, 3000} {
		sum += mixedPairs[(zeroIdx+p)%len(mixedPairs)].Num
	}

	return sum
}
