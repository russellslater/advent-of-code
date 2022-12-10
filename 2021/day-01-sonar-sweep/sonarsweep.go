package main

import (
	"fmt"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

func SonarSweep(readings []string, measurer func([]string, int) (int, bool)) (depthIncreases int) {
	depthIncreases = 0

	prevSum := 0
	isFirstPass := true

	for i := 0; i < len(readings); i++ {
		currentSum, ok := measurer(readings, i)

		if !ok {
			break
		}

		if !isFirstPass && (currentSum > prevSum) {
			depthIncreases++
		}

		isFirstPass = false
		prevSum = currentSum
	}

	return
}

func OneMeasurementSum(readings []string, index int) (sum int, ok bool) {
	return MeasurementSum(readings, 1, index)
}

func ThreeMeasurementSum(readings []string, index int) (sum int, ok bool) {
	return MeasurementSum(readings, 3, index)
}

func MeasurementSum(readings []string, measureSize int, index int) (sum int, ok bool) {
	sum = 0
	ok = true

	maxIndex := index + measureSize

	if maxIndex > len(readings) {
		ok = false
		return
	}

	for i := index; i < maxIndex; i++ {
		currentNum, _ := strconv.Atoi(readings[i])
		sum += currentNum
	}

	return
}

func main() {
	inputLines := util.LoadInput("./2021/day-01-sonar-sweep/input.txt")

	oneMeasuredepthIncreases := SonarSweep(inputLines, OneMeasurementSum)

	fmt.Printf("Depth increases: %d\n", oneMeasuredepthIncreases)

	threeMeasureDepthIncreases := SonarSweep(inputLines, ThreeMeasurementSum)

	fmt.Printf("Depth increases with three-measurement window: %d\n", threeMeasureDepthIncreases)
}
