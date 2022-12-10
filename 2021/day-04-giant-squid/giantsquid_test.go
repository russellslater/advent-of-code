package main

import (
	"reflect"
	"testing"
)

func TestGetBingoCalls(t *testing.T) {
	var tests = []struct {
		bingoCallLine  string
		wantBingoCalls []int
	}{
		{"", []int{}},
		{"1", []int{1}},
		{"30,35,8,2,39", []int{30, 35, 8, 2, 39}},
	}

	for _, test := range tests {
		ansBingoCalls := GetBingoCalls(test.bingoCallLine)

		if !reflect.DeepEqual(ansBingoCalls, test.wantBingoCalls) {
			t.Errorf("Result should be %v, instead %v for [%v]", test.wantBingoCalls, ansBingoCalls, test)
		}
	}
}

func TestGetBingoBoards(t *testing.T) {
	var tests = []struct {
		lines      []string
		wantBoards [][][]int
	}{
		{[]string{
			"22 13 17 11  0",
			" 8  2 23  4 24",
			"21  9 14 16  7",
			" 6 10  3 18  5",
			" 1 12 20 15 19",
			"",
			" 3 15  0  2 22",
			" 9 18 13 17  5",
			"19  8  7 25 23",
			"20 11 10 24  4",
			"14 21 16 12  6",
		},
			[][][]int{{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			}, {
				{3, 15, 0, 2, 22},
				{9, 18, 13, 17, 5},
				{19, 8, 7, 25, 23},
				{20, 11, 10, 24, 4},
				{14, 21, 16, 12, 6},
			}},
		},
	}

	for _, test := range tests {
		ansBoards := GetBingoBoards(test.lines)

		if !reflect.DeepEqual(ansBoards, test.wantBoards) {
			t.Errorf("Result should be %d, instead %d for [%v]", test.wantBoards, ansBoards, test)
		}
	}
}

func TestIsWinningBoard(t *testing.T) {
	var tests = []struct {
		markedNums   [][]bool
		wantIsWinner bool
	}{
		{[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{true, true, true, true, true},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}, true},
		{[][]bool{
			{false, false, true, false, false},
			{false, false, true, false, false},
			{true, false, true, true, true},
			{false, false, true, false, false},
			{false, false, true, false, false},
		}, true},
		{[][]bool{
			{false, true, true, false, false},
			{true, false, false, true, true},
			{true, false, true, true, true},
			{true, true, true, true, false},
			{false, true, true, true, true},
		}, false},
	}

	for _, test := range tests {
		ansIsWinner := IsWinningBoard(test.markedNums)

		if ansIsWinner != test.wantIsWinner {
			t.Errorf("Result should be %t, instead %t for [%v]", test.wantIsWinner, ansIsWinner, test)
		}
	}
}

func TestPlayBoard(t *testing.T) {
	var tests = []struct {
		board        [][]int
		calls        []int
		wantIsWinner bool
		wantScore    int
		wantDraws    int
	}{
		{[][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		}, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 19, 60}, true, 2736, 13},
		{[][]int{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		}, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 19, 60}, true, 4512, 12},
	}

	for _, test := range tests {
		ansIsWinner, ansScore, ansDraws := PlayBoard(test.board, test.calls)
		if ansIsWinner != test.wantIsWinner || ansScore != test.wantScore || ansDraws != test.wantDraws {
			t.Errorf("Result should be [%t, %d, %d], instead [%t, %d, %d] for [%v]",
				test.wantIsWinner, test.wantScore, test.wantDraws, ansIsWinner, ansScore, ansDraws, test)
		}
	}
}
