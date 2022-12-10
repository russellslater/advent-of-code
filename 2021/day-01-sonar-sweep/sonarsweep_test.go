package main

import (
	"testing"
)

func TestSonarSweep(t *testing.T) {
	var tests = []struct {
		readings []string
		want     int
	}{
		{[]string{"100", "101", "102", "100", "105"}, 3},
		{[]string{"500", "400", "300", "200", "100"}, 0},
		{[]string{"1", "2", "3", "4", "5", "6"}, 5},
		{[]string{"199", "200", "208", "210", "200", "207", "240", "260", "269", "263"}, 7},
	}

	for _, test := range tests {
		ans := SonarSweep(test.readings, OneMeasurementSum)
		if ans != test.want {
			t.Errorf("Number of depth increases should be %d, instead %d", test.want, ans)
		}
	}
}

func TestThreeMeasurementSum(t *testing.T) {
	var tests = []struct {
		readings []string
		index    int
		want     int
		wantOk   bool
	}{
		{[]string{"100"}, 0, 0, false},
		{[]string{}, 0, 0, false},
		{[]string{"100", "101", "102", "103", "104", "105"}, 0, 303, true},
		{[]string{"100", "101", "102", "103", "104", "105"}, 1, 306, true},
		{[]string{"100", "101", "102", "103", "104", "105"}, 2, 309, true},
		{[]string{"100", "101", "102", "103", "104", "105"}, 3, 312, true},
		{[]string{"100", "101", "102", "103", "104", "105"}, 4, 0, false},
	}

	for _, test := range tests {
		ans, ok := ThreeMeasurementSum(test.readings, test.index)
		if ans != test.want {
			t.Errorf("Three measurement sum should be %d, instead %d", test.want, ans)
		}
		if ok != test.wantOk {
			t.Errorf("Three measurement sum OK check should be %t, instead %t, for expected sum of %d", test.wantOk, ok, test.want)
		}
	}
}

func TestSonarSweepWithThreeMeasurementSums(t *testing.T) {
	var tests = []struct {
		readings []string
		want     int
	}{
		{[]string{"199", "200", "208", "210", "200", "207", "240", "260", "269", "263"}, 5},
		{[]string{"1", "2", "3", "4"}, 1},
		{[]string{"4", "3", "2", "1"}, 0},
		{[]string{}, 0},
		{[]string{"1"}, 0},
		{[]string{"1", "1", "1", "1", "1", "1"}, 0},
		{[]string{"100", "200", "100", "200", "100", "200"}, 2},
	}

	for _, test := range tests {
		ans := SonarSweep(test.readings, ThreeMeasurementSum)
		if ans != test.want {
			t.Errorf("Number of depth increases should be %d, instead %d", test.want, ans)
		}
	}
}
