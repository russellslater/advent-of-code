package main

import (
	"reflect"
	"testing"
)

func TestProcessSegmentsForDanger(t *testing.T) {
	var tests = []struct {
		lines            []string
		includeDiagonals bool
		wantOverlapCount int
	}{
		{[]string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}, false, 5},
	}
	for _, test := range tests {
		ansOverlapCount := ProcessSegmentsForDanger(test.lines, test.includeDiagonals)

		if ansOverlapCount != test.wantOverlapCount {
			t.Errorf("Result should be %d, instead %d for [%v]", test.wantOverlapCount, ansOverlapCount, test)
		}
	}
}

func TestCountSegmentOverlaps(t *testing.T) {
	var tests = []struct {
		oceanFloor       [][]int
		wantOverlapCount int
	}{
		{[][]int{
			{2, 1, 1, 2},
			{1, 0, 1, 2},
			{1, 0, 0, 1},
			{1, 0, 0, 1},
		}, 3},
	}
	for _, test := range tests {
		ansOverlapCount := CountSegmentOverlaps(test.oceanFloor)

		if ansOverlapCount != test.wantOverlapCount {
			t.Errorf("Result should be %d, instead %d for [%v]", test.wantOverlapCount, ansOverlapCount, test)
		}
	}
}

func TestPlotSegments(t *testing.T) {
	var tests = []struct {
		segments         []Segment
		maxX             int
		maxY             int
		includeDiagonals bool
		wantOceanFloor   [][]int
	}{
		{[]Segment{
			{ // vertical
				x1: 3,
				y1: 0,
				x2: 3,
				y2: 3,
			},
			{ // vertical
				x1: 0,
				y1: 0,
				x2: 0,
				y2: 3,
			},
			{ // horizontal
				x1: 0,
				y1: 0,
				x2: 3,
				y2: 0,
			},
			{ // horizontal
				x1: 2,
				y1: 1,
				x2: 3,
				y2: 1,
			},
		}, 3, 3, false,
			[][]int{
				{2, 1, 1, 2},
				{1, 0, 1, 2},
				{1, 0, 0, 1},
				{1, 0, 0, 1},
			}},
	}
	for _, test := range tests {
		ansOceanFloor := PlotSegments(test.segments, test.maxX, test.maxY, test.includeDiagonals)

		if !reflect.DeepEqual(ansOceanFloor, test.wantOceanFloor) {
			t.Errorf("Result should be %v, instead %v for [%v]", test.wantOceanFloor, ansOceanFloor, test)
		}
	}
}

func TestGetCoveredCoordinates(t *testing.T) {
	var tests = []struct {
		segment    Segment
		wantCoords [][]int
	}{
		{Segment{ // vertical
			x1: 0,
			y1: 0,
			x2: 0,
			y2: 1,
		}, [][]int{{0, 0}, {0, 1}}},
		{Segment{ // vertical
			x1: 0,
			y1: 0,
			x2: 0,
			y2: 5,
		}, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}},
		{Segment{ // vertical
			x1: 10,
			y1: 7,
			x2: 10,
			y2: 5,
		}, [][]int{{10, 5}, {10, 6}, {10, 7}}},
		{Segment{ // horizontal
			x1: 5,
			y1: 13,
			x2: 3,
			y2: 13,
		}, [][]int{{3, 13}, {4, 13}, {5, 13}}},
		{Segment{ // horizontal
			x1: 5,
			y1: 3,
			x2: 10,
			y2: 3,
		}, [][]int{{5, 3}, {6, 3}, {7, 3}, {8, 3}, {9, 3}, {10, 3}}},
		{Segment{ // diagonal
			x1: 0,
			y1: 0,
			x2: 2,
			y2: 2,
		}, [][]int{{0, 0}, {1, 1}, {2, 2}}},
		{Segment{ // diagonal
			x1: 1,
			y1: 1,
			x2: 3,
			y2: 3,
		}, [][]int{{1, 1}, {2, 2}, {3, 3}}},
		{Segment{ // diagonal
			x1: 9,
			y1: 7,
			x2: 7,
			y2: 9,
		}, [][]int{{9, 7}, {8, 8}, {7, 9}}},
		{Segment{ // diagonal
			x1: 5,
			y1: 11,
			x2: 11,
			y2: 5,
		}, [][]int{{5, 11}, {6, 10}, {7, 9}, {8, 8}, {9, 7}, {10, 6}, {11, 5}}},
	}

	for _, test := range tests {
		ansCoords := test.segment.getCoveredCoordinates()

		if !reflect.DeepEqual(ansCoords, test.wantCoords) {
			t.Errorf("Result should be %v, instead %v for [%v]", test.wantCoords, ansCoords, test)
		}
	}
}

func TestIsDiagonal(t *testing.T) {
	var tests = []struct {
		segment        Segment
		wantIsDiagonal bool
	}{
		{Segment{
			x1: 0,
			y1: 0,
			x2: 2,
			y2: 2,
		}, true},
		{Segment{
			x1: 1,
			y1: 1,
			x2: 3,
			y2: 3,
		}, true},
		{Segment{
			x1: 9,
			y1: 7,
			x2: 7,
			y2: 9,
		}, true},
		{Segment{ // diagonal
			x1: 5,
			y1: 11,
			x2: 11,
			y2: 5,
		}, true},
	}

	for _, test := range tests {
		ansIsDiagonal := test.segment.isDiagonal()

		if ansIsDiagonal != test.wantIsDiagonal {
			t.Errorf("Result should be %t, instead %t for [%v]", test.wantIsDiagonal, ansIsDiagonal, test)
		}
	}
}

func TestIsVertical(t *testing.T) {
	var tests = []struct {
		segment        Segment
		wantIsVertical bool
	}{
		{Segment{
			x1: 0,
			y1: 5,
			x2: 0,
			y2: 23,
		}, true},
		{Segment{
			x1: 0,
			y1: 23,
			x2: 0,
			y2: 5,
		}, true},
		{Segment{
			x1: 57,
			y1: 173,
			x2: 57,
			y2: 400,
		}, true},
		{Segment{
			x1: 296,
			y1: 172,
			x2: 646,
			y2: 522,
		}, false},
	}

	for _, test := range tests {
		ansIsVertical := test.segment.isVertical()

		if ansIsVertical != test.wantIsVertical {
			t.Errorf("Result should be %t, instead %t for [%v]", test.wantIsVertical, ansIsVertical, test)
		}
	}
}

func TestIsHorizontal(t *testing.T) {
	var tests = []struct {
		segment          Segment
		wantIsHorizontal bool
	}{
		{Segment{
			x1: 0,
			y1: 5,
			x2: 10,
			y2: 5,
		}, true},
		{Segment{
			x1: 173,
			y1: 57,
			x2: 400,
			y2: 57,
		}, true},
		{Segment{
			x1: 296,
			y1: 172,
			x2: 646,
			y2: 522,
		}, false},
	}

	for _, test := range tests {
		ansIsHorizontal := test.segment.isHorizontal()

		if ansIsHorizontal != test.wantIsHorizontal {
			t.Errorf("Result should be %t, instead %t for [%v]", test.wantIsHorizontal, ansIsHorizontal, test)
		}
	}
}

func TestCreateSegment(t *testing.T) {
	var tests = []struct {
		line        string
		wantSegment Segment
	}{
		{"296,172 -> 646,522",
			Segment{
				x1: 296,
				y1: 172,
				x2: 646,
				y2: 522,
			}},
		{"61,427 -> 395,427",
			Segment{
				x1: 61,
				y1: 427,
				x2: 395,
				y2: 427,
			}},
	}

	for _, test := range tests {
		ansSegment := CreateSegment(test.line)

		if ansSegment != test.wantSegment {
			t.Errorf("Result should be %v, instead %v for [%v]", test.wantSegment, ansSegment, test)
		}
	}
}
