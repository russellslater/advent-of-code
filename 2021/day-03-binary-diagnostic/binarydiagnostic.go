package main

import (
	"fmt"
	"strconv"

	"github.com/russellslater/advent-of-code/internal/util"
)

func FilterByBit(readings []string, bitPosition int, bitValue string) (filteredReadings []string) {
	for i := 0; i < len(readings); i++ {
		if readings[i][bitPosition] == []byte(bitValue)[0] {
			filteredReadings = append(filteredReadings, readings[i])
		}
	}

	return
}

func FindMostPopularBit(readings []string, bitPosition int) (bitValue string) {
	bitValue = FindPopularBit(readings, bitPosition, true)
	return
}

func FindLeastPopularBit(readings []string, bitPosition int) (bitValue string) {
	bitValue = FindPopularBit(readings, bitPosition, false)
	return
}

func FindPopularBit(readings []string, bitPosition int, isMostPopular bool) (bitValue string) {
	freq := 0

	for i := 0; i < len(readings); i++ {
		if readings[i][bitPosition] == '1' {
			freq++
		}
	}

	if ((freq*2 >= len(readings)) && isMostPopular) || (freq*2 < len(readings) && !isMostPopular) {
		bitValue = "1"
	} else {
		bitValue = "0"
	}

	return
}

func OxygenGeneratorRating(readings []string) (rating int) {
	rating = FindRating(readings, FindMostPopularBit)
	return
}

func CO2ScrubberRating(readings []string) (rating int) {
	rating = FindRating(readings, FindLeastPopularBit)
	return
}

func FindRating(readings []string, popularBitFinder func([]string, int) string) (rating int) {
	filteredReadings := readings

	// loop through every bit position
	for i := 0; i < len(readings[0]); i++ {
		bitValue := popularBitFinder(filteredReadings, i)
		filteredReadings = FilterByBit(filteredReadings, i, bitValue)

		if len(filteredReadings) == 1 {
			parsedRating, _ := strconv.ParseInt(filteredReadings[0], 2, 64)
			rating = int(parsedRating)
			break
		}
	}

	return
}

func LifeSupportRating(readings []string) (lifeSupportRating int) {
	oRating := OxygenGeneratorRating(readings)
	cRating := CO2ScrubberRating(readings)

	lifeSupportRating = oRating * cRating

	return lifeSupportRating
}

func DiagnosticReport(readings []string) (gammaRate int, epsilonRate int) {
	counters := make([]int, len(readings[0]))

	for i := 0; i < len(readings); i++ {
		for pos, char := range readings[i] {
			if char == '1' {
				counters[pos] += 1
			}
		}
	}

	count := len(readings)
	binaryGammaStr := ""
	binaryEpsilonStr := ""

	for _, freq := range counters {
		if freq*2 > count {
			binaryGammaStr += "1"
			binaryEpsilonStr += "0"
		} else {
			binaryGammaStr += "0"
			binaryEpsilonStr += "1"
		}
	}

	binaryResult, _ := strconv.ParseInt(binaryGammaStr, 2, 64)
	epsilonResult, _ := strconv.ParseInt(binaryEpsilonStr, 2, 64)

	gammaRate, epsilonRate = int(binaryResult), int(epsilonResult)

	return
}

func main() {
	inputLines := util.LoadInput("./2021/day-03-binary-diagnostic/input.txt")

	gammaRate, epsilonRate := DiagnosticReport(inputLines)
	product := gammaRate * epsilonRate

	fmt.Printf("Diagnostic Report - Gamma Rate: %d, Epsilon Rate: %d, Product: %d\n", gammaRate, epsilonRate, product)

	lifeSupportRating := LifeSupportRating(inputLines)

	fmt.Printf("Life Support Rating: %d\n", lifeSupportRating)
}
