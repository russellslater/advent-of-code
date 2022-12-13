package main

import (
	"encoding/json"
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-13-distress-signal/input.txt"
	rawPacketPairs := getTransformedInput(filename)

	sumIndices, packetsReceived := calcSumIndicesOfRightOrderPairs(rawPacketPairs)
	fmt.Printf("Part One Answer: %d\n", sumIndices)

	decoderKey := calcDecoderKey(packetsReceived)
	fmt.Printf("Part Two Answer: %d\n", decoderKey)
}

func getTransformedInput(filename string) [][]string {
	pairs := [][]string{}
	pairs = append(pairs, []string{})

	lines := util.LoadInput(filename)

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			pairs = append(pairs, []string{})
			continue
		} else {
			pairs[len(pairs)-1] = append(pairs[len(pairs)-1], lines[i])
		}
	}

	return pairs
}

func calcSumIndicesOfRightOrderPairs(rawPacketPairs [][]string) (int, []any) {
	sumIndices := 0
	packetsReceived := []any{}

	for i, pair := range rawPacketPairs {
		left := parsePacket(pair[0])
		right := parsePacket(pair[1])

		if compare(left, right) < 0 {
			sumIndices += i + 1
		}

		packetsReceived = append(packetsReceived, left, right)
	}

	return sumIndices, packetsReceived
}

func calcDecoderKey(packets []any) int {
	// Add divider packets
	dividerOne, dividerTwo := parsePacket("[[2]]"), parsePacket("[[6]]")
	packets = append(packets, dividerOne, dividerTwo)

	// Sort packets in ascending order
	for i := 0; i < len(packets); i++ {
		for j := i + 1; j < len(packets); j++ {
			if compare(packets[i], packets[j]) > 0 {
				packets[i], packets[j] = packets[j], packets[i]
			}
		}
	}

	product := 1

	// Decoder key is the product of the indices of the divider packets
	for i, packet := range packets {
		if compare(packet, dividerOne) == 0 || compare(packet, dividerTwo) == 0 {
			product *= i + 1
		}
	}

	return product
}

func parsePacket(packetData string) any {
	var packet any
	json.Unmarshal([]byte(packetData), &packet)
	return packet
}

func compare(left, right any) int {
	// If both parts of the pair are integers, compare them as integers.
	// Note: json.Unmarshal stores numbers as float64
	firstVal, okFirst := left.(float64)
	secondVal, okSecond := right.(float64)

	if okFirst && okSecond {
		return int(firstVal) - int(secondVal)
	}

	// If both parts of the pair are lists, compare them as lists.
	// If one value is an integer, convert that value to a list with one element.

	var leftList []any
	var rightList []any

	switch left.(type) {
	case []any, []float64:
		leftList = left.([]any)
	case float64:
		leftList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		rightList = right.([]any)
	case float64:
		rightList = []any{right}
	}

	for i := range leftList {
		if i >= len(rightList) {
			// Right side ran out of items, so inputs are not in the right order
			return 1
		}

		// Return first non-zero comparison result
		if res := compare(leftList[i], rightList[i]); res != 0 {
			return res
		}
	}

	if len(leftList) == len(rightList) {
		// Elements are equal, lengths are equal
		return 0
	}

	// Left side ran out of items, so inputs are in the right order
	return -1
}
