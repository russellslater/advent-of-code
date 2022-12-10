package main

import (
	"reflect"
	"testing"
)

func TestDiagnosticReport(t *testing.T) {
	var tests = []struct {
		readings    []string
		wantGamma   int
		wantEpsilon int
	}{
		{[]string{"101011010010", "111011011000"}, 2768, 1327},
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010"}, 22, 9},
	}

	for _, test := range tests {
		ansGamma, ansEpsilon := DiagnosticReport(test.readings)
		if ansGamma != test.wantGamma || ansEpsilon != test.wantEpsilon {
			t.Errorf("Command should be [%d, %d], instead [%d, %d]", test.wantGamma, test.wantEpsilon, ansGamma, ansEpsilon)
		}
	}
}

func TestLifeSupportRating(t *testing.T) {
	var tests = []struct {
		readings              []string
		wantLifeSupportRating int
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010"}, 0},
	}

	for _, test := range tests {
		ansLifeSupportReading := LifeSupportRating(test.readings)
		if ansLifeSupportReading != test.wantLifeSupportRating {
			t.Errorf("Command should be %d, instead %d", test.wantLifeSupportRating, ansLifeSupportReading)
		}
	}
}

func TestFilterByBit(t *testing.T) {
	var tests = []struct {
		readings             []string
		bitPosition          int
		bitValue             string
		wantFilteredReadings []string
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010"}, 0, "1", []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"},
		},
		{[]string{
			"11110",
			"10110",
			"10111",
			"10101",
			"11100",
			"10000",
			"11001"}, 1, "0", []string{"10110", "10111", "10101", "10000"},
		},
	}

	for _, test := range tests {
		ansFilteredReadings := FilterByBit(test.readings, test.bitPosition, test.bitValue)
		if !reflect.DeepEqual(ansFilteredReadings, test.wantFilteredReadings) {
			t.Errorf("Command should be %v, instead %v", test.wantFilteredReadings, ansFilteredReadings)
		}
	}
}

func TestFindMostPopularBit(t *testing.T) {
	var tests = []struct {
		readings     []string
		bitPosition  int
		wantBitValue string
	}{
		{[]string{"1"}, 0, "1"},
		{[]string{"0"}, 0, "0"},
		{[]string{"1", "1"}, 0, "1"},
		{[]string{"1", "0"}, 0, "1"}, // if equal, defaults to "1"
		{[]string{"1", "0", "1"}, 0, "1"},
		{[]string{"1", "0", "0"}, 0, "0"},
		{[]string{
			"11110",
			"10110",
			"10111",
			"10101",
			"11100",
			"10000",
			"11001"}, 1, "0"},
	}

	for _, test := range tests {
		ansBitValue := FindMostPopularBit(test.readings, test.bitPosition)
		if ansBitValue != test.wantBitValue {
			t.Errorf("Command should be %v, instead %v when %v", test.wantBitValue, ansBitValue, test)
		}
	}
}

func TestOxygenGeneratorRating(t *testing.T) {
	var tests = []struct {
		readings   []string
		wantRating int
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010"}, 23},
	}

	for _, test := range tests {
		ansRating := OxygenGeneratorRating(test.readings)
		if ansRating != test.wantRating {
			t.Errorf("Command should be %v, instead %v when %v", test.wantRating, ansRating, test)
		}
	}
}

func TestCO2ScrubberRating(t *testing.T) {
	var tests = []struct {
		readings   []string
		wantRating int
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010"}, 10},
	}

	for _, test := range tests {
		ansRating := CO2ScrubberRating(test.readings)
		if ansRating != test.wantRating {
			t.Errorf("Command should be %v, instead %v when %v", test.wantRating, ansRating, test)
		}
	}
}
