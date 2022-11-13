package main

import (
	"fmt"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	nums := getNums("2020/day-01-report-repair/input.txt")
	sumTarget := 2020

	pairProduct := FindPairSumProduct(sumTarget, nums)

	fmt.Printf("Pair Product = %d\n", pairProduct)

	trioProduct := FindTrioSumProduct(sumTarget, nums)

	fmt.Printf("Trio Product = %d\n", trioProduct)
}

func getNums(filename string) []int {
	input := []int{}
	for _, line := range util.LoadInput(filename) {
		num, _ := strconv.Atoi(line)
		input = append(input, num)
	}
	return input
}

func FindPairSumProduct(sumTarget int, nums []int) int {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n+nums[j] == sumTarget {
				fmt.Printf("Pair found [%d, %d]\n", n, nums[j])
				return n * nums[j]
			}
		}
	}

	return 0
}

func FindTrioSumProduct(sumTarget int, nums []int) int {
	for i, n := range nums {
		res := FindPairSumProduct(sumTarget-n, nums[i+1:])
		if res != 0 {
			fmt.Printf("Trio found [%d]\n", n)
			return n * res
		}
	}

	return 0
}
